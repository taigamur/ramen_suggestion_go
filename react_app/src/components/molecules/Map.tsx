import { memo, useEffect, useState } from "react"
import { GoogleMap, LoadScript, Marker } from "@react-google-maps/api";
import axios from "axios";

type Geocode = {
    lat: number;
    lng: number;
}

type Props = {
    place: string;
}

const containerStyle = {
    width: "100%",
    height: "30vh",
};

export const Map = memo((props: Props) => {
    const { place } = props;

    const [ geocode, setGeocode ] = useState<Geocode>()

    const getGeocode = (place: string) => {
        var position: Geocode = {
            lat: 0,
            lng: 0,
        };
        axios.get('https://maps.googleapis.com/maps/api/geocode/json', 
            {params: 
                {address: place ,
                key: process.env.REACT_APP_GOOGLE! }
            })
        .then((res) => {
            position.lat = res.data.results[0].geometry.location.lat;
            position.lng = res.data.results[0].geometry.location.lng;
            setGeocode(position)
        })
        .catch(() => {
            console.log("error");
        });
    }

    useEffect(() => getGeocode(place),[place])

    return (
        <>
        <LoadScript googleMapsApiKey={process.env.REACT_APP_GOOGLE!}>
            <GoogleMap mapContainerStyle={containerStyle} center={geocode!} zoom={15} >
                <Marker position={geocode!} />
            </GoogleMap>
        </LoadScript>
        </>
      );
})

