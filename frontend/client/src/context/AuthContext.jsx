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
    const fetchUser = async (token) => {
      try {
        const userResponse = await userService.getCurrent(token);
        dispatch({
          type: ACTIONS.LOGIN,
          payload: {
            token: token,
            user: userResponse.user,
          },
        });
      } catch (err) {
        console.error(err);
      }
    };
    const token = localStorage.getItem("token");
    if (token) {
      fetchUser(token);
    }
  }, []);

  const login = async (email, password) => {
    const loginResponse = await userService.login(email, password);
    const userResponse = await userService.getCurrent(loginResponse.token);

    dispatch({
      type: ACTIONS.LOGIN,
      payload: {
        token: loginResponse.token,
        user: userResponse.user,
      },
    });
  };

  const socialLogin = async () => {
    const response = await userService.getToken();
    const userResponse = await userService.getCurrent(response.token);

    dispatch({
      type: ACTIONS.LOGIN,
      payload: {
        token: response.token,
        user: userResponse.user,
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
    <AuthContext.Provider value={{ ...state, login, socialLogin, logout }}>
      {children}
    </AuthContext.Provider>
  );
};

// eslint-disable-next-line react-refresh/only-export-components
export const useAuthContext = () => React.useContext(AuthContext);
export default AuthProvider;
