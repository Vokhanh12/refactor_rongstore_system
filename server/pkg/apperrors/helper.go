package apperrors

func MergeErrors(errs ...*AppError) *AppError {
	var details []AppErrorDetail

	for _, err := range errs {
		if err == nil {
			continue
		}
		details = append(details, err.ErrorDetails...)
	}

	if len(details) == 0 {
		return nil
	}

	return New(errs[0], WithAppendErrorDetails(details))
}

func copyError(src AppError) *AppError {
	dst := src

	if src.Tags != nil {
		dst.Tags = append([]string(nil), src.Tags...)
	}

	dst.Data = copyDataMap(src.Data)
	dst.ErrorDetails = copyDetails(src.ErrorDetails)

	return &dst
}

func copyDataMap(src map[string]interface{}) map[string]interface{} {
	if src == nil {
		return nil
	}

	dst := make(map[string]interface{}, len(src))
	for k, v := range src {
		dst[k] = v
	}
	return dst
}

func copyDetails(src []AppErrorDetail) []AppErrorDetail {
	if src == nil {
		return nil
	}

	dst := make([]AppErrorDetail, len(src))
	copy(dst, src)
	return dst
}
