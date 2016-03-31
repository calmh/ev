var pid = 0;
var start = 0;
var end = 0;
var chartSeriesMap = [
    {chart: '#chart-cpu', series: [{i:0, t:'area'}, {i:1, t:'area'}], name: "CPU Usage (%)"},
    {chart: '#chart-ram', series: [{i:3, t:'area'}, {i:2, t:'area'}, {i:6, t:'line'}], name: "RAM Usage (bytes)"},
    {chart: '#chart-gr', series: [{i:4, t:'area'}], name: "Goroutines"},
    {chart: '#chart-gc', series: [{i:5, t:'area'}], name: "Allocation Rate (bytes/s)"},
];

Highcharts.setOptions({global: {useUTC: false}});

$.getJSON('/data', function(data) {
    $.each(data, function(i, choice) {
        $('#pids')
            .append($('<option>', { value: choice })
                .text(choice));
    });
    pid = data[0];
    load();
});

function selectPID() {
    pid = $('#pids').find(':selected').text();
    refresh();
}

function syncExtremes(e) {
    var thisChart = this.chart;

    if (e.trigger !== 'syncExtremes') { // Prevent feedback loop
        Highcharts.each(Highcharts.charts, function(chart) {
            if (chart !== thisChart) {
                if (chart.xAxis[0].setExtremes) { // It is null while updating
                    chart.xAxis[0].setExtremes(e.min, e.max, undefined, false, { trigger: 'syncExtremes' });
                }
            }
        });
    }
}

function afterSetExtremes(e) {
    if (e.trigger === 'zoom') { // Prevent feedback loop
        start = Math.round(e.userMin);
        end = Math.round(e.userMax);
        console.log("efter",e.trigger)
        refresh()
    }
}

function refresh() {
    console.log("refresh")
    $.getJSON('/data?pid=' + pid + '&start=' + start + '&end=' + end, function(data) {
        for (var i = 0; i < chartSeriesMap.length; i++) {
            var dsName = chartSeriesMap[i].name;
            var chartID = chartSeriesMap[i].chart;
            var chart = $(chartID).highcharts();
            for (var j = 0; j < chartSeriesMap[i].series.length; j++) {
                var seriesIdx = chartSeriesMap[i].series[j].i;
                var seriesData = Highcharts.map(data.Dataset.Time, function(t, k) {
                    return [t, data.Dataset.Series[seriesIdx].Values[k]];
                });
                chart.series[j].setData(seriesData);
            }
        }

        var res = "";
        for (var i = 0; i < data.Lines.length; i++) {
            res = res + data.Lines[i].Line + "\n";
        }
        $('#log').html(res)
    });
}

function periodicRefresh() {
    console.log("periodicRefresh")
    if (!start && !end) {
        refresh();
    }
}

function load() {
    console.log("load")
    $.getJSON('/data?pid=' + pid, function(data) {
        setInterval(periodicRefresh, 2500);

        for (var i = 0; i < chartSeriesMap.length; i++) {
            var dsName = chartSeriesMap[i].name;
            var chartID = chartSeriesMap[i].chart;
            var series = [];
            for (var j = 0; j < chartSeriesMap[i].series.length; j++) {
                var seriesIdx = chartSeriesMap[i].series[j].i;
                var seriesData = Highcharts.map(data.Dataset.Time, function(t, k) {
                    return [t, data.Dataset.Series[seriesIdx].Values[k]];
                });
                var seriesType = chartSeriesMap[i].series[j].t;
                series.push({
                    type: seriesType,
                    name: data.Dataset.Series[seriesIdx].Name,
                    data: seriesData,
                    marker: {
                        enabled: false
                    }
                });
            }

            $(chartID).highcharts({
                chart: {
                    marginLeft: 80, // Keep all charts left aligned
                    spacingTop: 20,
                    spacingBottom: 20,
                    zoomType: 'x',
                    animation: false,
                },
                title: {
                    text: dsName,
                    align: 'left',
                    margin: 0,
                    x: 30
                },
                credits: {
                    enabled: false
                },
                legend: {
                    enabled: false
                },

                xAxis: {
                    type: 'datetime',
                    crosshair: true,
                    events: {
                        setExtremes: syncExtremes,
                        afterSetExtremes: afterSetExtremes
                    }
                },
                yAxis: {
                    title: {
                        text: null
                    }
                },
                tooltip: {
                    enabled: false,
                    //pointFormat: '{point.y}',
                    //headerFormat: '',
                },
                plotOptions: {
                    series: {
                        animation: false,
                    }
                },
                series: series,
                exporting: {
                    enabled: false
                }
            });
        }
    });
}
