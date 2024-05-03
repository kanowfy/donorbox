import { Button } from "flowbite-react";
import { useNavigate } from "react-router-dom";
import { useAuthContext } from "../context/AuthContext";

const Login = () => {
  const { login } = useAuthContext();
  const navigate = useNavigate();

  async function submitLoginDetails() {
    try {
      await login("", "");
      navigate("/");
    } catch (err) {
      console.error(err);
    }
  }

  return <Button onClick={submitLoginDetails}>Login</Button>;
};

export default Login;
