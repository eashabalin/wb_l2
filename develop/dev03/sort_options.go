package dev03

import "errors"

type SortOptions struct {
	FileName            string
	OutputFileName      string
	Column              int
	Unique              bool
	Reverse             bool
	Numeric             bool
	Check               bool
	Month               bool
	IgnoreLeadingBlanks bool
	SISuffix            bool
}

func (o *SortOptions) Validate() error {
	if o.Month && o.Numeric {
		return errors.New("can't combine -M and -n")
	}
	if o.Numeric && o.SISuffix {
		return errors.New("can't combine -n and -h")
	}
	if o.SISuffix && o.Month {
		return errors.New("can't combine -h and -M")
	}
	return nil
}
