package ledger

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
}

// Column widths of the rendered table.
const (
	dateWidth        = 10
	descriptionWidth = 25
	changeWidth      = 13
)

const isoDateLayout = "2006-01-02"

type localeFormat struct {
	dateHeader   string
	descHeader   string
	changeHeader string
	dateLayout   string
	formatAmount func(symbol string, cents int) string
}

var locales = map[string]localeFormat{
	"en-US": {
		dateHeader:   "Date",
		descHeader:   "Description",
		changeHeader: "Change",
		dateLayout:   "01/02/2006",
		formatAmount: formatAmountUS,
	},
	"nl-NL": {
		dateHeader:   "Datum",
		descHeader:   "Omschrijving",
		changeHeader: "Verandering",
		dateLayout:   "02-01-2006",
		formatAmount: formatAmountNL,
	},
}

var currencySymbols = map[string]string{
	"USD": "$",
	"EUR": "€",
}

func FormatLedger(currency string, locale string, entries []Entry) (string, error) {
	format, ok := locales[locale]
	if !ok {
		return "", fmt.Errorf("unsupported locale %q", locale)
	}
	symbol, ok := currencySymbols[currency]
	if !ok {
		return "", fmt.Errorf("unsupported currency %q", currency)
	}

	var out strings.Builder
	out.WriteString(formatHeader(format))
	for _, entry := range sortedEntries(entries) {
		row, err := formatEntry(entry, format, symbol)
		if err != nil {
			return "", err
		}
		out.WriteString(row)
	}
	return out.String(), nil
}

func formatHeader(format localeFormat) string {
	return joinColumns(
		format.dateHeader,
		format.descHeader,
		format.changeHeader,
	)
}

func formatEntry(entry Entry, format localeFormat, symbol string) (string, error) {
	date, err := time.Parse(isoDateLayout, entry.Date)
	if err != nil {
		return "", fmt.Errorf("invalid date %q", entry.Date)
	}
	amount := format.formatAmount(symbol, entry.Change)
	return joinColumns(
		date.Format(format.dateLayout),
		truncate(entry.Description, descriptionWidth),
		padLeft(amount, changeWidth),
	), nil
}

func joinColumns(date, description, change string) string {
	return strings.Join([]string{
		padRight(date, dateWidth),
		padRight(description, descriptionWidth),
		padRight(change, changeWidth),
	}, " | ") + "\n"
}

// sortedEntries returns a copy ordered by date, description, then change.
func sortedEntries(entries []Entry) []Entry {
	sorted := make([]Entry, len(entries))
	copy(sorted, entries)
	sort.SliceStable(sorted, func(i, j int) bool {
		return isBefore(sorted[i], sorted[j])
	})
	return sorted
}

func isBefore(a, b Entry) bool {
	if a.Date != b.Date {
		return a.Date < b.Date
	}
	if a.Description != b.Description {
		return a.Description < b.Description
	}
	return a.Change < b.Change
}

// formatAmountUS renders "$1,234.56", negatives in parentheses.
func formatAmountUS(symbol string, cents int) string {
	negative, units, fraction := splitAmount(cents)
	amount := symbol + groupDigits(units, ",") + "." + fraction
	if negative {
		return "(" + amount + ")"
	}
	return amount + " "
}

// formatAmountNL renders "$ 1.234,56", negatives with a leading minus.
func formatAmountNL(symbol string, cents int) string {
	negative, units, fraction := splitAmount(cents)
	sign := ""
	if negative {
		sign = "-"
	}
	return symbol + " " + sign + groupDigits(units, ".") + "," + fraction + " "
}

func splitAmount(cents int) (negative bool, units int, fraction string) {
	negative = cents < 0
	if negative {
		cents = -cents
	}
	return negative, cents / 100, fmt.Sprintf("%02d", cents%100)
}

// groupDigits inserts sep between groups of three digits: 1234567 -> "1,234,567".
func groupDigits(n int, sep string) string {
	digits := strconv.Itoa(n)
	head := len(digits) % 3
	if head == 0 {
		head = 3
	}
	groups := []string{digits[:head]}
	for i := head; i < len(digits); i += 3 {
		groups = append(groups, digits[i:i+3])
	}
	return strings.Join(groups, sep)
}

// truncate shortens s to width characters, marking the cut with "...".
func truncate(s string, width int) string {
	if utf8.RuneCountInString(s) <= width {
		return s
	}
	return string([]rune(s)[:width-3]) + "..."
}

func padRight(s string, width int) string {
	return s + padding(s, width)
}

func padLeft(s string, width int) string {
	return padding(s, width) + s
}

func padding(s string, width int) string {
	missing := width - utf8.RuneCountInString(s)
	if missing <= 0 {
		return ""
	}
	return strings.Repeat(" ", missing)
}
