package quartz_test

import (
	"testing"
	"time"

	"github.com/reugn/go-quartz/quartz"
)

func TestCronExpression1(t *testing.T) {
	prev := int64(1555351200000000000)
	result := ""
	cronTrigger, err := quartz.NewCronTrigger("10/20 15 14 5-10 * ? *")
	cronTrigger.Description()
	if err != nil {
		t.Fatal(err)
	} else {
		result, _ = iterate(prev, cronTrigger, 1000)
	}
	assertEqual(t, result, "Fri Dec 8 14:15:10 2023")
}

func TestCronExpression2(t *testing.T) {
	prev := int64(1555351200000000000)
	result := ""
	cronTrigger, err := quartz.NewCronTrigger("* 5,7,9 14-16 * * ? *")
	if err != nil {
		t.Fatal(err)
	} else {
		result, _ = iterate(prev, cronTrigger, 1000)
	}
	assertEqual(t, result, "Mon Aug 5 14:05:00 2019")
}

func TestCronExpression3(t *testing.T) {
	prev := int64(1555351200000000000)
	result := ""
	cronTrigger, err := quartz.NewCronTrigger("* 5,7,9 14/2 * * Wed,Sat *")
	if err != nil {
		t.Fatal(err)
	} else {
		result, _ = iterate(prev, cronTrigger, 1000)
	}
	assertEqual(t, result, "Sat Dec 7 14:05:00 2019")
}

func TestCronExpression4(t *testing.T) {
	expression := "0 5,7 14 1 * Sun *"
	_, err := quartz.NewCronTrigger(expression)
	if err == nil {
		t.Fatalf("%s should fail", expression)
	}
}

func TestCronExpression5(t *testing.T) {
	prev := int64(1555351200000000000)
	result := ""
	cronTrigger, err := quartz.NewCronTrigger("* * * * * ? *")
	if err != nil {
		t.Fatal(err)
	} else {
		result, _ = iterate(prev, cronTrigger, 1000)
	}
	assertEqual(t, result, "Mon Apr 15 18:16:40 2019")
}

func TestCronExpression6(t *testing.T) {
	prev := int64(1555351200000000000)
	result := ""
	cronTrigger, err := quartz.NewCronTrigger("* * 14/2 * * Mon/3 *")
	if err != nil {
		t.Fatal(err)
	} else {
		result, _ = iterate(prev, cronTrigger, 1000)
	}
	assertEqual(t, result, "Mon Mar 15 18:00:00 2021")
}

func TestCronExpression7(t *testing.T) {
	prev := int64(1555351200000000000)
	result := ""
	cronTrigger, err := quartz.NewCronTrigger("* 5-9 14/2 * * 0-2 *")
	if err != nil {
		t.Fatal(err)
	} else {
		result, _ = iterate(prev, cronTrigger, 1000)
	}
	assertEqual(t, result, "Tue Jul 16 16:09:00 2019")
}

func TestCronYearly(t *testing.T) {
	prev := int64(1555351200000000000)
	result := ""
	cronTrigger, err := quartz.NewCronTrigger("@yearly")
	if err != nil {
		t.Fatal(err)
	} else {
		result, _ = iterate(prev, cronTrigger, 100)
	}
	assertEqual(t, result, "Sun Jan 1 00:00:00 2119")
}

func TestCronMonthly(t *testing.T) {
	prev := int64(1555351200000000000)
	result := ""
	cronTrigger, err := quartz.NewCronTrigger("@monthly")
	if err != nil {
		t.Fatal(err)
	} else {
		result, _ = iterate(prev, cronTrigger, 100)
	}
	assertEqual(t, result, "Sun Aug 1 00:00:00 2027")
}

func TestCronWeekly(t *testing.T) {
	prev := int64(1555351200000000000)
	result := ""
	cronTrigger, err := quartz.NewCronTrigger("@weekly")
	if err != nil {
		t.Fatal(err)
	} else {
		result, _ = iterate(prev, cronTrigger, 100)
	}
	assertEqual(t, result, "Mon Mar 15 00:00:00 2021")
}

func TestCronDaily(t *testing.T) {
	prev := int64(1555351200000000000)
	result := ""
	cronTrigger, err := quartz.NewCronTrigger("@daily")
	if err != nil {
		t.Fatal(err)
	} else {
		result, _ = iterate(prev, cronTrigger, 1000)
	}
	assertEqual(t, result, "Sun Jan 9 00:00:00 2022")
}

func TestCronHourly(t *testing.T) {
	prev := int64(1555351200000000000)
	result := ""
	cronTrigger, err := quartz.NewCronTrigger("@hourly")
	if err != nil {
		t.Fatal(err)
	} else {
		result, _ = iterate(prev, cronTrigger, 1000)
	}
	assertEqual(t, result, "Wed May 29 06:00:00 2019")
}

var readDateLayout = "Mon Jan 2 15:04:05 2006"

func iterate(prev int64, cronTrigger *quartz.CronTrigger, iterations int) (string, error) {
	var err error
	for i := 0; i < iterations; i++ {
		prev, err = cronTrigger.NextFireTime(prev)
		// fmt.Println(time.Unix(prev/int64(time.Second), 0).UTC().Format(readDateLayout))
		if err != nil {
			return "", err
		}
	}
	return time.Unix(prev/int64(time.Second), 0).UTC().Format(readDateLayout), nil
}
