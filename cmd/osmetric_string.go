// Code generated by "stringer -type=osMetric -trimprefix=osMetric os-instrumented.go"; DO NOT EDIT.

package cmd

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[osMetricRemoveAll-0]
	_ = x[osMetricMkdirAll-1]
	_ = x[osMetricRename-2]
	_ = x[osMetricOpenFile-3]
	_ = x[osMetricOpen-4]
	_ = x[osMetricOpenFileDirectIO-5]
	_ = x[osMetricLstat-6]
	_ = x[osMetricRemove-7]
	_ = x[osMetricStat-8]
	_ = x[osMetricAccess-9]
	_ = x[osMetricCreate-10]
	_ = x[osMetricReadDirent-11]
	_ = x[osMetricFdatasync-12]
	_ = x[osMetricSync-13]
	_ = x[osMetricLast-14]
}

const _osMetric_name = "RemoveAllMkdirAllRenameOpenFileOpenOpenFileDirectIOLstatRemoveStatAccessCreateReadDirentFdatasyncSyncLast"

var _osMetric_index = [...]uint8{0, 9, 17, 23, 31, 35, 51, 56, 62, 66, 72, 78, 88, 97, 101, 105}

func (i osMetric) String() string {
	if i >= osMetric(len(_osMetric_index)-1) {
		return "osMetric(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _osMetric_name[_osMetric_index[i]:_osMetric_index[i+1]]
}
