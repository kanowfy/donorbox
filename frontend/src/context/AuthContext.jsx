import React from "react";
import axios from "axios";
import userService from "../services/user";
import { AUTH_ACTIONS as ACTIONS } from "../constants";

const AuthContext = React.createContext();

const authReducer = (state, action) => {
  switch (action.type) {
    case ACTIONS.LOGIN:
      axios.defaults.headers.common[
        "Authorization"
      ] = `Bearer ${action.payload.token}`;
      localStorage.setItem("token", action.payload.token);
      return {
        ...state,
        token: action.payload.token,
        user: action.payload.user,
      };
    case ACTIONS.LOGOUT:
      delete axios.defaults.headers.common["Authorization"];
      localStorage.removeItem("token");
      return { ...state, token: null, user: null };
    default:
      throw new Error(`Unhandled action type: ${action.type}`);
  }
};

// eslint-disable-next-line react/prop-types
const AuthProvider = ({ children }) => {
  const [state, dispatch] = React.useReducer(authReducer, {
    token: null,
    user: null,
  });

  React.useEffect(() => {
    const token = localStorage.getItem("token");
    const fetchUser = async () => {
      try {
        if (token) {
          const user = await userService.getCurrentUser(
            localStorage.getItem("token")
          );
          console.log(user);
          dispatch({
            type: ACTIONS.LOGIN,
            payload: {
              token: token,
              user: user,
            },
          });
        }
      } catch (err) {
        console.error(err);
      }
    };

    fetchUser();
  }, []);

  const login = async (email, password) => {
    const response = await userService.login(email, password);
    const user = await userService.getCurrentUser(response.token);

    dispatch({
      type: ACTIONS.LOGIN,
      payload: {
        token: response.token,
        user: user,
      },
    });
  };

  const logout = async () => {
    dispatch({
      type: ACTIONS.LOGOUT,
      payload: null,
    });
  };

  return (
    <AuthContext.Provider value={{ ...state, login, logout }}>
      {children}
    </AuthContext.Provider>
  );
};

// eslint-disable-next-line react-refresh/only-export-components
export const useAuthContext = () => React.useContext(AuthContext);
export default AuthProvider;
