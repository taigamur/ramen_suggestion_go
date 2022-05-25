import { memo, useEffect, useState } from "react"
import { GoogleMap, LoadScript, Marker } from "@react-google-maps/api";
import axios from "axios";
import { Link } from "@chakra-ui/react"
import { ExternalLinkIcon } from "@chakra-ui/icons";

type Geocode = {
    lat: number;
    lng: number;
}

type Props = {
    address: string;
    name: string;
}

const containerStyle = {
    width: "100%",
    height: "30vh",
};

export const Map = memo((props: Props) => {
    const { address, name } = props;

    const [ geocode, setGeocode ] = useState<Geocode>()
    const [ url, setURL ] = useState<string>("");

    const getGeocode = (address: string) => {
        setURL('https://www.google.com/maps/search/?api=1&query=' + name + "+つくば")
        console.log(url)
        var position: Geocode = {
            lat: 0,
            lng: 0,
        };
        axios.get('https://maps.googleapis.com/maps/api/geocode/json', 
            {params: 
                {address: address ,
                key: process.env.REACT_APP_GOOGLE! }
            })
        .then((res) => {
            position.lat = res.data.results[0].geometry.location.lat;
            position.lng = res.data.results[0].geometry.location.lng;
            console.log(position)
            setGeocode(position)
        })
        .catch(() => {
            console.log("error");
        });
    }

    useEffect(() => getGeocode(address),[address])

    return (
        <>
        <Link href={url} color='blue.600' isExternal>
            Google Mapで開く <ExternalLinkIcon mx='2px' />
        </Link>
        <LoadScript googleMapsApiKey={process.env.REACT_APP_GOOGLE!}>
            <GoogleMap mapContainerStyle={containerStyle} center={geocode!} zoom={15} >
                <Marker position={geocode!} />
            </GoogleMap>
        </LoadScript>
        </>
      );
})

