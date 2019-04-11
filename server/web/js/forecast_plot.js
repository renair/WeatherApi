$forecastModal = $("#forecastModal");

function displayForecast(dates, temperatures) {
	var Temperature = {
		x: dates,
		y: temperatures,
		line: {
			color: "red",
			shape: "spline"
		},
		name: 'temperature',
		type: 'scatter',
	};
	
	var data = [Temperature];
	
	var layout = {
		showlegend: true
	};
	
	Plotly.newPlot('forecastPlot', data, layout, {
		showSendToCloud: false,
		scrollZoom: true
	});
	$forecastModal.modal('show');
}