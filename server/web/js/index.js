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
	
	if(!localStorage.getItem('isFirstTime')){
		$("#helpModal").modal('show');
		localStorage.setItem('isFirstTime','true')
	}
	
	if (navigator.geolocation) {
		navigator.geolocation.getCurrentPosition(function(position){
			moveMapTo(position.coords);
		});
	}
});