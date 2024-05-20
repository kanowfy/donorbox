import { useEffect } from "react";
import { useAuthContext } from "../context/AuthContext";
import { useNavigate } from "react-router-dom";

const SocialLogin = () => {
  const { socialLogin } = useAuthContext();
  const navigate = useNavigate();

  useEffect(() => {
    const login = async () => {
      try {
        await socialLogin();
        navigate("/");
      } catch (err) {
        console.error(err);
      }
    };
    login();
  }, [navigate, socialLogin]);

  return <div>Login successful! Redirecting to home...</div>;
};

export default SocialLogin;
