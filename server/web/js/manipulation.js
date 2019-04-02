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

function weatherCardContent(weather) {
	$card = $weatherCard.clone();
	$card.find(".location-name").text(weather.location.locationName);
	$card.find(".weather-temperature").text(weather.values.temperature);
	$card.find(".weather-humidity").text(weather.values.humidity + "%");
	$card.find(".weather-pressure").text(weather.values.pressure + "mm");
//	const updateTime = new Date(weather.date*1000 + Date.getTimezoneOffset()*60*1000).toGMTString();
	$card.find(".weather-last-update").text(new Date(weather.date*1000).toLocaleString());
	$card.show();
	return $card.html();
}