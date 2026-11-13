export const API_URL = import.meta.env.API_URL ||
        (typeof process !== "undefined"
            ? process.env.API_URL
            : undefined);
export const APP_ORIGIN = import.meta.env.APP_ORIGIN ||
        (typeof process !== "undefined"
            ? process.env.APP_ORIGIN
            : undefined);