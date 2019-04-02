let MAP;
let onMarkerClick;

function initMap() {
	MAP = new google.maps.Map(document.getElementById('map'), {
		center: {lat: 50.42, lng: 30.63},
		zoom: 12
	});
	
	MAP.addListener('click', function(data){
		openModal(data.latLng.lat(), data.latLng.lng());
	});
}

function displayMarker(location) {
	var marker = new google.maps.Marker({
		position: {'lat': location.latitude, 'lng': location.longitude},
		map: MAP,
		title: location.locatinName
	});
	marker.addListener('click', function() {
		MAP.panTo(marker.position);
		if(onMarkerClick) onMarkerClick(marker, location.id);
	});
}

function showInfoWindow(marker, content) {
	var infowindow = new google.maps.InfoWindow({
		content: content
	});
	infowindow.open(MAP, marker);
}

