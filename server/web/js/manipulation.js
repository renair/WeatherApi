const $modal = $('#newLocationModal');
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
	
	apiCreateLocation(lat, lon, name, () => {
		displayMarker(lat, lon, name);
		$modal.find("#locationNameInput").val("");
	});
	$modal.modal('hide');
}