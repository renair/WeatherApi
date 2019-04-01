let map;

function initMap() {
	console.log(document.getElementById('map'));
	map = new google.maps.Map(document.getElementById('map'), {
		center: {lat: 50.42, lng: 30.63},
		zoom: 12
	});
}

initMap();