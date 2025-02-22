import { useForm } from "react-hook-form";
import { useNavigate } from "react-router-dom";
import { useAuthContext } from "../../context/AuthContext";
import { Button } from "@tremor/react";

const Login = () => {
  const { login } = useAuthContext();
  const navigate = useNavigate();
  const {
    register,
    formState: { errors },
    setError,
    handleSubmit,
  } = useForm();

  function onSubmit(data) {
    const loginAccount = async (data) => {
      try {
        await login(data.email, data.password);
        navigate("/");
      } catch (err) {
        console.error(err);
        setError("email", {
          type: "manual",
          message: "Invalid email or password",
        });
      }
    };

    loginAccount(data);
  }

  return (
    <section className="bg-gradient-to-tr from-green-500 to-yellow-200">
      <div className="flex flex-col items-center justify-center px-6 py-8 mx-auto md:h-screen lg:py-0">
        <a
          href="#"
          className="flex items-center mb-6 text-2xl font-semibold text-gray-800 tracking-tight"
        >
          <img className="w-8 h-8 mr-2" src="/logo.svg" alt="logo" />
          Donorbox
        </a>
        <div className="w-full bg-white rounded-lg shadow md:mt-0 sm:max-w-md xl:p-0">
          <div className="p-6 space-y-4 md:space-y-6 sm:p-8">
            <h1 className="text-xl font-bold leading-tight tracking-tight text-gray-900 md:text-2xl">
              Sign in to your account
            </h1>
            <form
              onSubmit={handleSubmit(onSubmit)}
              className="space-y-4 md:space-y-6"
            >
              <div>
                {errors.email?.type === "manual" && (
                  <p className="text-red-600 text-sm">
                    Invalid email or password
                  </p>
                )}
                <label className="block mb-2 text-sm font-medium text-gray-900">
                  Email
                </label>
                <input
                  {...register("email", {
                    required: true,
                  })}
                  type="email"
                  className="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5"
                  placeholder="Your Email"
                />
                {errors.email?.type === "required" && (
                  <p className="text-red-600 text-sm">Email is required</p>
                )}
              </div>
              <div>
                <label className="block mb-2 text-sm font-medium text-gray-900">
                  Password
                </label>
                <input
                  {...register("password", {
                    required: true,
                  })}
                  type="password"
                  placeholder="••••••••"
                  className="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5"
                />
                {errors.password?.type === "required" && (
                  <p className="text-red-600 text-sm">Password is required</p>
                )}
              </div>

              <Button type="submit" color="teal" className="w-full">
                Sign in
              </Button>
            </form>
          </div>
        </div>
      </div>
    </section>
  );
};

export default Login;
