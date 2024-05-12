import { useNavigate } from "react-router-dom";
import { useAuthContext } from "../context/AuthContext";
import { useEffect } from "react";

const Profile = () => {
  const { user } = useAuthContext();
  const navigate = useNavigate();

  useEffect(() => {
    if (!user) {
      navigate("/");
    }
    console.log(user);
  }, [user, navigate]);

  return <div>Profile</div>;
};

export default Profile;
