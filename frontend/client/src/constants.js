export const BASE_URL = import.meta.env.VITE_BACKEND_URL;
export const SERVE_URL = import.meta.env.VITE_SERVE_URL;
export const AUTH_ACTIONS = {
    LOGIN: "login",
    LOGOUT: "logout",
    LOADED: "loaded"
}

export const CategoryIndexMap = {
    Medical: 1,
    Emergency: 2,
    Education: 3,
    Animals: 4,
    Competition: 5,
    Event: 6,
    Environment: 7,
    Travel: 8,
    Business: 9,
};
