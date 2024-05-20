import { Button } from "flowbite-react";
import { Link } from "react-router-dom";

const RegisterSuccess = () => {
  return (
    <section className="bg-gradient-to-tr from-yellow-200 to-green-500 min-h-screen flex items-center justify-center">
      <div className="bg-white rounded-lg shadow-xl flex flex-col items-center py-10 px-5">
        <div>
          <img src="/success.svg" height={72} width={72} />
        </div>
        <div className="font-bold text-2xl mt-5 mb-3">Register successful!</div>
        <div className="tracking-tight mb-5">
          Please follow the confirmation email we have sent to you to complete
          the registration process
        </div>
        <Link to="/">
          <Button color="green">Go to Homepage</Button>
        </Link>
      </div>
    </section>
  );
};

export default RegisterSuccess;
