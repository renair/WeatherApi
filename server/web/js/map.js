let MAP;

function initMap() {
	console.log(document.getElementById('map'));
	MAP = new google.maps.Map(document.getElementById('map'), {
		center: {lat: 50.42, lng: 30.63},
		zoom: 12
	});
	
	MAP.addListener('click', function(data){
		openModal(data.latLng.lat(), data.latLng.lng());
	});
}

