import React from "react";
import authService from "../services/auth";
import { AUTH_ACTIONS as ACTIONS } from "../constants";

const AuthContext = React.createContext();

const authReducer = (state, action) => {
  switch (action.type) {
    case ACTIONS.LOGIN:
      localStorage.setItem("token", action.payload.token);
      return {
        ...state,
        token: action.payload.token,
        user: action.payload.user,
      };
    case ACTIONS.LOGOUT:
      localStorage.removeItem("token");
      return { ...state, token: null, user: null };
    default:
      throw new Error(`Unhandled action type: ${action.type}`);
  }
};

// eslint-disable-next-line react/prop-types
const AuthProvider = ({ children }) => {
  const [state, dispatch] = React.useReducer(authReducer, {
    token: localStorage.getItem("token"),
    user: null,
  });

  React.useEffect(() => {
    const fetchUser = async (token) => {
      try {
        const userResponse = await authService.getCurrent(token);
        console.log(userResponse.user);
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
    const loginResponse = await authService.login(email, password);
    const userResponse = await authService.getCurrent(loginResponse.token);

    dispatch({
      type: ACTIONS.LOGIN,
      payload: {
        token: loginResponse.token,
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
    <AuthContext.Provider value={{ ...state, login, logout }}>
      {children}
    </AuthContext.Provider>
  );
};

// eslint-disable-next-line react-refresh/only-export-components
export const useAuthContext = () => React.useContext(AuthContext);
export default AuthProvider;
