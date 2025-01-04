import { Button } from "flowbite-react";
import { useForm } from "react-hook-form";
import { Link, useNavigate } from "react-router-dom";
import userService from "../../services/user";

const Register = () => {
  const {
    register,
    formState: { errors },
    handleSubmit,
    getValues,
  } = useForm();
  const navigate = useNavigate();

  function onSubmit(data) {
    const registerAccount = async (data) => {
      try {
        await userService.register(data);
        navigate("/register/success");
      } catch (err) {
        console.error(err);
      }
    };
    registerAccount(data);
  }

  return (
    <section className="bg-gradient-to-tr from-sky-200 to-sky-400">
      <div className="flex flex-col items-center justify-center px-6 py-8 mx-auto md:h-screen lg:py-0">
        <Link
          to="/"
          className="flex items-center mb-6 text-2xl font-semibold text-gray-800 tracking-tight"
        >
          <img className="w-8 h-8 mr-2" src="/logo.svg" alt="logo" />
          Donorbox
        </Link>
        <div className="w-full bg-white rounded-lg shadow-xl md:mt-0 sm:max-w-md xl:p-0">
          <div className="p-6 space-y-4 md:space-y-6 sm:p-8">
            <h1 className="text-xl font-bold leading-tight tracking-tight text-gray-900 md:text-2xl">
              Create an account
            </h1>
            <form
              onSubmit={handleSubmit(onSubmit)}
              className="space-y-4 md:space-y-6"
            >
              <div>
                <label className="block mb-2 text-sm font-medium text-gray-900">
                  Email <span className="text-red-700">*</span>
                </label>
                <input
                  {...register("email", {
                    required: true,
                    maxLength: 50,
                    pattern: { value: /\S+@\S+\.\S+/ },
                  })}
                  type="email"
                  className="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5"
                  placeholder="Your Email"
                />
                {errors.email?.type === "required" && (
                  <p className="text-red-600 text-sm">Email is required</p>
                )}
                {errors.email?.type === "maxLength" && (
                  <p className="text-red-600 text-sm">Max email length is 50</p>
                )}
                {errors.email?.type === "pattern" && (
                  <p className="text-red-600 text-sm">Invalid email format</p>
                )}
              </div>
              <div className="flex gap-4">
                <div>
                  <label className="block mb-2 text-sm font-medium text-gray-900">
                    First name <span className="text-red-700">*</span>
                  </label>
                  <input
                    {...register("first_name", {
                      required: true,
                      minLength: 2,
                      maxLength: 50,
                    })}
                    type="text"
                    className="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5"
                    placeholder="John"
                  />
                  {errors.firstName?.type === "required" && (
                    <p className="text-red-600 text-sm">
                      First name is required
                    </p>
                  )}
                  {errors.firstName?.type === "minLength" && (
                    <p className="text-red-600 text-sm">Minimum length is 2</p>
                  )}
                  {errors.firstName?.type === "maxLength" && (
                    <p className="text-red-600 text-sm">Maximum length is 50</p>
                  )}
                </div>
                <div>
                  <label className="block mb-2 text-sm font-medium text-gray-900">
                    Last name <span className="text-red-700">*</span>
                  </label>
                  <input
                    {...register("last_name", {
                      required: true,
                      minLength: 2,
                      maxLength: 50,
                    })}
                    type="text"
                    className="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5"
                    placeholder="Doe"
                  />
                  {errors.lastName?.type === "required" && (
                    <p className="text-red-600 text-sm">
                      Last name is required
                    </p>
                  )}
                  {errors.lastName?.type === "minLength" && (
                    <p className="text-red-600 text-sm">Minimum length is 2</p>
                  )}
                  {errors.lastName?.type === "maxLength" && (
                    <p className="text-red-600 text-sm">Maximum length is 50</p>
                  )}
                </div>
              </div>
              <div>
                <label className="block mb-2 text-sm font-medium text-gray-900">
                  Password <span className="text-red-700">*</span>
                </label>
                <input
                  {...register("password", {
                    required: true,
                    minLength: 8,
                    maxLength: 50,
                  })}
                  type="password"
                  placeholder="••••••••"
                  className="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5"
                />
                {errors.password?.type === "required" && (
                  <p className="text-red-600 text-sm">Password is required</p>
                )}
                {errors.password?.type === "minLength" && (
                  <p className="text-red-600 text-sm">Minimum length is 8</p>
                )}
                {errors.password?.type === "maxLength" && (
                  <p className="text-red-600 text-sm">Maximum length is 50</p>
                )}
              </div>
              <div>
                <label className="block mb-2 text-sm font-medium text-gray-900">
                  Confirm password <span className="text-red-700">*</span>
                </label>
                <input
                  {...register("repassword", {
                    required: true,
                    validate: (value) => {
                      const { password } = getValues();
                      return password === value || "Password must match";
                    },
                  })}
                  type="password"
                  placeholder="••••••••"
                  className="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5"
                />
                {errors.repassword?.type === "required" && (
                  <p className="text-red-600 text-sm">
                    Please confirm your password
                  </p>
                )}
                {errors.repassword?.type === "validate" && (
                  <p className="text-red-600 text-sm">
                    Password does not match
                  </p>
                )}
              </div>
              <Button
                gradientDuoTone="greenToBlue"
                className="w-full font-semibold"
                type="submit"
              >
                Sign up
              </Button>
              <p className="text-sm font-normal text-gray-500">
                Already have an account?{" "}
                <a
                  href="/login"
                  className="font-medium text-primary-600 hover:underline"
                >
                  Login here
                </a>
              </p>
            </form>
          </div>
        </div>
      </div>
    </section>
  );
};

export default Register;
