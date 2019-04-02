$(document).ready(function(){
	initMap();
	apiGetRegionLocations(50, 30, 10000, (data) => {
		data.locationsInRegion.forEach((loc) => displayMarker(loc.latitude, loc.longitude, loc.locationName));
	});
});