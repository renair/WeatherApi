const gqlQueries = `
mutation CreateLocation($lng: Float!, $lat: Float!, $name: String) {
  createLocation (input: {longitude: $lng, latitude: $lat, locationName: $name}) {
    id
  }
}

query GetLocationById($id: Int! = 1) {
  location(id: $id) {
    longitude
    latitude
    locationName
    id
  }
}

query GetLocationsByCoords($lng: Float!, $lat: Float!, $radius: Float!) {
  locationsInRegion(longitude: $lng, latitude: $lat, radius: $radius) {
    longitude
    latitude
    locationName
    id
  }
}

query GetRegionWeather($lng: Float!, $lat: Float!, $radius: Float!) {
  weatherInRegion(longitude: $lng, latitude: $lat, radius: $radius) {
    values {
      temperature
      pressure
      humidity
    }
    wind {
      speed
      direction
    }
    location { 
    	longitude
      latitude
    }
    cloud {
      isSnow
      isRain
    }
    date
  }
}

query GetWeatherById($id: Int!) {
  weatherInLocation(locationId: $id) {
    values {
      temperature
      pressure
      humidity
    }
    location { 
    	longitude
      latitude
    }
    cloud {
      isSnow
      isRain
    }
    date
  }
}`;

const gqlEndpoint = window.location.origin + "/query";

function qraphQlQuery(operationName, args, callback, errCallback) {
	const reqData = {
		"query": gqlQueries,
		"operationName": operationName,
		"variables": args
	}
	$.ajax({
        'type': 'POST',
        'url': gqlEndpoint,
        'contentType': 'application/json',
        'data': JSON.stringify(reqData),
        'dataType': 'json'
    }).done((data) => {
		if(data.errors) {
			if(errCallback) errCallback(data.errors);
			return;
		}
		if(callback) callback(data.data);
	})
	.fail((err) => {
		console.log(err);
		errCallback(err);
	});
}

function apiGetRegionLocations(lat, lng, rad, callback, errCallback) {
	const reqArgs = {
		'lat': lat,
		'lng': lng,
		'radius': rad
	};
	qraphQlQuery("GetLocationsByCoords", reqArgs, callback, errCallback);
}

function apiCreateLocation(lat, lng, name, callback, errCallback) {
	const reqArgs = {
		'lat': lat,
		'lng': lng,
		'name': name
	};
	qraphQlQuery("CreateLocation", reqArgs, callback, errCallback);
}