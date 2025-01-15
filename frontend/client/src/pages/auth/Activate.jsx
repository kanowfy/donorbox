import { useEffect, useState } from "react";
import { Link, useNavigate, useSearchParams } from "react-router-dom";
import userService from "../../services/user";
import { Button } from "flowbite-react";

const Activate = () => {
  const [errorMessage, setErrorMessage] = useState(null);
  const [searchParams] = useSearchParams();
  const token = searchParams.get("token");
  const navigate = useNavigate();
  useEffect(() => {
    const activateAccount = async (token) => {
      try {
        await userService.verify(token);
        setTimeout(() => {
          navigate("/login");
        }, 3000);
      } catch (err) {
        console.error(err);
        setErrorMessage(err.response.data.error);
      }
    };

    activateAccount(token);
  }, [token, navigate]);

  return (
    <section className="bg-gradient-to-tr from-yellow-200 to-green-500 min-h-screen flex items-center justify-center">
      <div className="bg-white rounded-lg shadow-xl flex flex-col items-center py-10 px-5">
        {errorMessage ? (
          <>
            <div>
              <img src="/fail.svg" height={72} width={72} />
            </div>
            <div className="font-bold text-2xl mt-5 mb-3">
              Activation failed!
            </div>
            <div className="tracking-tight mb-5">{errorMessage}</div>
            <Link to="/login">
              <Button color="green">Go to Login</Button>
            </Link>
          </>
        ) : (
          <>
            <div>
              <img src="/success.svg" height={72} width={72} />
            </div>
            <div className="font-bold text-2xl mt-5 mb-3">
              Activation successful!
            </div>
            <div className="tracking-tight mb-5">
              You will be navigated to the login page...
            </div>
          </>
        )}
      </div>
    </section>
  );
};

export default Activate;
