const $modal = $('#newLocationModal');
const $weatherCard = $('.weather-card');

$weatherCard.hide();
$modal.find("#createLocationButton").click(createLocation);

function openModal(lat, lon) {
	$modal.modal('show');
	$modal.find("#locationLatitudeInput").val(lat);
	$modal.find("#locationLongitudeInput").val(lon);
}

function createLocation() {
	const name = $modal.find("#locationNameInput").val();
	const lat = parseFloat($modal.find("#locationLatitudeInput").val());
	const lon = parseFloat($modal.find("#locationLongitudeInput").val());
	
	if(name == "") {
		name = null;
	}
	
	if(isNaN(lat) || isNaN(lon)) {
		alert("Coordinates shold be correct float value!");
		return;
	}
	
	apiCreateLocation(lat, lon, name, (data) => {
		displayMarker(data.createLocation);
		$modal.find("#locationNameInput").val("");
	});
	$modal.modal('hide');
}

function weatherCard(weather) {
	$card = $weatherCard.clone();
	$card.find(".location-name").text(weather.location.locationName);
	$forecastModal.find(".location-name").text(weather.location.locationName);
	$card.find(".weather-temperature").text(weather.values.temperature);
	$card.find(".weather-humidity").text(weather.values.humidity + "%");
	$card.find(".weather-pressure").text(weather.values.pressure + "mm");
	$card.find(".weather-last-update").text(new Date(weather.date*1000).toLocaleString());
	$card.find(".get-forecast").click(getForecastClicked(weather.location.id));
	$card.show();
	return $card;
}

function getForecastClicked(locationId) {
	return function(){
		apiGetForecast(locationId, (data) => {
			var dates = data.weatherInLocation.forecast.map((item) => new Date(item.date*1000).toLocaleString());
			var values = data.weatherInLocation.forecast.map((item) => item.values.temperature);
			displayForecast(dates, values);
		},
		(err)=>{
			console.log("Error in getting forecast!");
			console.log(err);
		})
	};
}