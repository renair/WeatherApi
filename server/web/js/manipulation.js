const $modal = $('#newLocationModal');

function openModal(lat, lon) {
	$modal.modal('show');
	$modal.find("#locationLatitudeInput").val(lat);
	$modal.find("#locationLongitudeInput").val(lon);
}