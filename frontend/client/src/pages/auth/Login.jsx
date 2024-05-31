import { Button } from "flowbite-react";
import { useForm } from "react-hook-form";
import { Link, useNavigate, useSearchParams } from "react-router-dom";
import { useAuthContext } from "../../context/AuthContext";
import { BASE_URL } from "../../constants";

const Login = () => {
  const { login } = useAuthContext();
  const navigate = useNavigate();
  const [params] = useSearchParams();
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
        const redirect = params.get("redirect");
        if (redirect) {
          navigate(`/${redirect}`);
        } else {
          navigate("/");
        }
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

  function handleSSO() {
    window.location.href = `${BASE_URL}/users/auth/google`;
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

              <div className="flex items-center">
                <div className="h-px bg-gray-300 w-full"></div>
                <div className="text-gray-500 mx-2">or</div>
                <div className="h-px bg-gray-300 w-full"></div>
              </div>

              <Button
                color="light"
                className="w-full"
                outline
                onClick={handleSSO}
              >
                <img src="/google.svg" className="h-6 mr-1" />
                <span className="block">Sign in with Google</span>
              </Button>

              <div className="flex items-center justify-between">
                <div className="flex items-start">
                  <div className="flex items-center h-5">
                    <input
                      id="remember"
                      aria-describedby="remember"
                      type="checkbox"
                      className="w-4 h-4 border border-gray-300 rounded bg-gray-50 focus:ring-3 focus:ring-primary-300"
                      required=""
                    />
                  </div>
                  <div className="ml-3 text-sm">
                    <label className="text-gray-600">Remember me</label>
                  </div>
                </div>
                <Link
                  to="/password/forgot"
                  className="text-sm font-medium text-primary-600 hover:underline"
                >
                  Forgot password?
                </Link>
              </div>
              <Button
                gradientDuoTone="greenToBlue"
                className="w-full"
                type="submit"
              >
                Sign in
              </Button>
              <p className="text-sm font-normal text-gray-500">
                Don’t have an account yet?{" "}
                <a
                  href="/register"
                  className="font-medium text-primary-600 hover:underline"
                >
                  Sign up
                </a>
              </p>
            </form>
          </div>
        </div>
      </div>
    </section>
  );
};

export default Login;
