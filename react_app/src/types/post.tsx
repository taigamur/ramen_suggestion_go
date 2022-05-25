import { Place } from "../types/place"

export type Post = {
    id: number;
    place_id: number;
    username: string;
    value: number;
    date: string;
    place: Place;
}