import { useCallback, useState } from "react"
import { Place } from "../types/place"

type Props = {
    id: number;
    places: Array<Place>
    onOpen: () => void;
}

export const useSelectPlace = () => {
    const [ selectedPlace, setSelectedPlace ] = useState<Place | null>(null);

    const onSelectPlace = useCallback((props: Props) => {
        const { id, places, onOpen } = props;
        const targetPlace = places.find((place) => place.id === id)
        setSelectedPlace(targetPlace ?? null)
        onOpen()
    }, [])
    return { selectedPlace, onSelectPlace }
}