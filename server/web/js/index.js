$(document).ready(function(){
	initMap();
	apiGetRegionLocations(50, 30, 10000, (data) => {
		data.locationsInRegion.forEach((loc) => displayMarker(loc));
	});
	
	onMarkerClick = function(marker, id) {
		console.log("Loading weather for location with id: " + id);
		apiGetWeather(id, function(data){
			const content = weatherCardContent(data.weatherInLocation);
			showInfoWindow(marker, content);
		});
	}
});