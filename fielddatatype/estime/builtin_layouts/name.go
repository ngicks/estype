package builtinlayouts

const (
	EpochMillis                        = "epoch_millis"
	EpochSecond                        = "epoch_second"
	DateOptionalTime                   = "date_optional_time"
	StrictDateOptionalTime             = "strict_date_optional_time"
	StrictDateOptionalTimeNanos        = "strict_date_optional_time_nanos"
	BasicDate                          = "basic_date"
	BasicDateTime                      = "basic_date_time"
	BasicDateTimeNoMillis              = "basic_date_time_no_millis"
	BasicOrdinalDate                   = "basic_ordinal_date"
	BasicOrdinalDateTime               = "basic_ordinal_date_time"
	BasicOrdinalDateTimeNoMillis       = "basic_ordinal_date_time_no_millis"
	BasicTime                          = "basic_time"
	BasicTimeNoMillis                  = "basic_time_no_millis"
	BasicTTime                         = "basic_t_time"
	BasicTTimeNoMillis                 = "basic_t_time_no_millis"
	BasicWeekDate                      = "basic_week_date"
	StrictBasicWeekDate                = "strict_basic_week_date"
	BasicWeekDateTime                  = "basic_week_date_time"
	StrictBasicWeekDateTime            = "strict_basic_week_date_time"
	BasicWeekDateTimeNoMillis          = "basic_week_date_time_no_millis"
	StrictBasicWeekDateTimeNoMillis    = "strict_basic_week_date_time_no_millis"
	Date                               = "date"
	StrictDate                         = "strict_date"
	DateHour                           = "date_hour"
	StrictDateHour                     = "strict_date_hour"
	DateHourMinute                     = "date_hour_minute"
	StrictDateHourMinute               = "strict_date_hour_minute"
	DateHourMinuteSecond               = "date_hour_minute_second"
	StrictDateHourMinuteSecond         = "strict_date_hour_minute_second"
	DateHourMinuteSecondFraction       = "date_hour_minute_second_fraction"
	StrictDateHourMinuteSecondFraction = "strict_date_hour_minute_second_fraction"
	DateHourMinuteSecondMillis         = "date_hour_minute_second_millis"
	StrictDateHourMinuteSecondMillis   = "strict_date_hour_minute_second_millis"
	DateTime                           = "date_time"
	StrictDateTime                     = "strict_date_time"
	DateTimeNoMillis                   = "date_time_no_millis"
	StrictDateTimeNoMillis             = "strict_date_time_no_millis"
	Hour                               = "hour"
	StrictHour                         = "strict_hour"
	HourMinute                         = "hour_minute"
	StrictHourMinute                   = "strict_hour_minute"
	HourMinuteSecond                   = "hour_minute_second"
	StrictHourMinuteSecond             = "strict_hour_minute_second"
	HourMinuteSecondFraction           = "hour_minute_second_fraction"
	StrictHourMinuteSecondFraction     = "strict_hour_minute_second_fraction"
	HourMinuteSecondMillis             = "hour_minute_second_millis"
	StrictHourMinuteSecondMillis       = "strict_hour_minute_second_millis"
	OrdinalDate                        = "ordinal_date"
	StrictOrdinalDate                  = "strict_ordinal_date"
	OrdinalDateTime                    = "ordinal_date_time"
	StrictOrdinalDateTime              = "strict_ordinal_date_time"
	OrdinalDateTimeNoMillis            = "ordinal_date_time_no_millis"
	StrictOrdinalDateTimeNoMillis      = "strict_ordinal_date_time_no_millis"
	Time                               = "time"
	StrictTime                         = "strict_time"
	TimeNoMillis                       = "time_no_millis"
	StrictTimeNoMillis                 = "strict_time_no_millis"
	TTime                              = "t_time"
	StrictTTime                        = "strict_t_time"
	TTimeNoMillis                      = "t_time_no_millis"
	StrictTTimeNoMillis                = "strict_t_time_no_millis"
	WeekDate                           = "week_date"
	StrictWeekDate                     = "strict_week_date"
	WeekDateTime                       = "week_date_time"
	StrictWeekDateTime                 = "strict_week_date_time"
	WeekDateTimeNoMillis               = "week_date_time_no_millis"
	StrictWeekDateTimeNoMillis         = "strict_week_date_time_no_millis"
	Weekyear                           = "weekyear"
	StrictWeekyear                     = "strict_weekyear"
	WeekyearWeek                       = "weekyear_week"
	StrictWeekyearWeek                 = "strict_weekyear_week"
	WeekyearWeekDay                    = "weekyear_week_day"
	StrictWeekyearWeekDay              = "strict_weekyear_week_day"
	Year                               = "year"
	StrictYear                         = "strict_year"
	YearMonth                          = "year_month"
	StrictYearMonth                    = "strict_year_month"
	YearMonthDay                       = "year_month_day"
	StrictYearMonthDay                 = "strict_year_month_day"
)

var BuiltinFormats = []string{
	EpochMillis,
	EpochSecond,
	DateOptionalTime,
	StrictDateOptionalTime,
	StrictDateOptionalTimeNanos,
	BasicDate,
	BasicDateTime,
	BasicDateTimeNoMillis,
	BasicOrdinalDate,
	BasicOrdinalDateTime,
	BasicOrdinalDateTimeNoMillis,
	BasicTime,
	BasicTimeNoMillis,
	BasicTTime,
	BasicTTimeNoMillis,
	BasicWeekDate,
	StrictBasicWeekDate,
	BasicWeekDateTime,
	StrictBasicWeekDateTime,
	BasicWeekDateTimeNoMillis,
	StrictBasicWeekDateTimeNoMillis,
	Date,
	StrictDate,
	DateHour,
	StrictDateHour,
	DateHourMinute,
	StrictDateHourMinute,
	DateHourMinuteSecond,
	StrictDateHourMinuteSecond,
	DateHourMinuteSecondFraction,
	StrictDateHourMinuteSecondFraction,
	DateHourMinuteSecondMillis,
	StrictDateHourMinuteSecondMillis,
	DateTime,
	StrictDateTime,
	DateTimeNoMillis,
	StrictDateTimeNoMillis,
	Hour,
	StrictHour,
	HourMinute,
	StrictHourMinute,
	HourMinuteSecond,
	StrictHourMinuteSecond,
	HourMinuteSecondFraction,
	StrictHourMinuteSecondFraction,
	HourMinuteSecondMillis,
	StrictHourMinuteSecondMillis,
	OrdinalDate,
	StrictOrdinalDate,
	OrdinalDateTime,
	StrictOrdinalDateTime,
	OrdinalDateTimeNoMillis,
	StrictOrdinalDateTimeNoMillis,
	Time,
	StrictTime,
	TimeNoMillis,
	StrictTimeNoMillis,
	TTime,
	StrictTTime,
	TTimeNoMillis,
	StrictTTimeNoMillis,
	WeekDate,
	StrictWeekDate,
	WeekDateTime,
	StrictWeekDateTime,
	WeekDateTimeNoMillis,
	StrictWeekDateTimeNoMillis,
	Weekyear,
	StrictWeekyear,
	WeekyearWeek,
	StrictWeekyearWeek,
	WeekyearWeekDay,
	StrictWeekyearWeekDay,
	Year,
	StrictYear,
	YearMonth,
	StrictYearMonth,
	YearMonthDay,
	StrictYearMonthDay,
}
