import { Button } from "flowbite-react";
import { useForm } from "react-hook-form";
import user from "../../services/user";
import { useState } from "react";
import { Link } from "react-router-dom";

const ForgotPassword = () => {
  const [isSuccessful, setIsSuccessful] = useState(false);
  const {
    register,
    formState: { errors },
    handleSubmit,
  } = useForm();

  const onSubmit = async (data) => {
    try {
      await user.forgotPassword(data.email);
      setIsSuccessful(true);
    } catch (err) {
      console.error(err);
    }
  };

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
        <div className="w-full bg-white rounded-lg shadow md:mt-0 sm:max-w-md xl:p-0">
          <div className="p-6 space-y-4 md:space-y-6 sm:p-8">
            <form onSubmit={handleSubmit(onSubmit)} hidden={isSuccessful}>
              <div className="mb-5 font-medium text-lg tracking-tight leading-tight text-gray-800">
                Enter your email and we will send you a link to reset your
                password
              </div>
              <div>
                <label className="block mb-1 text-sm font-medium text-gray-700">
                  Email address
                </label>
                <input
                  {...register("email", {
                    required: true,
                    pattern: { value: /\S+@\S+\.\S+/ },
                  })}
                  type="email"
                  className="bg-gray-50 border border-gray-300 text-gray-900 rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-3"
                  placeholder="Your Email"
                />
                {errors.email?.type === "required" && (
                  <p className="text-red-600 text-sm">Email is required</p>
                )}
                {errors.email?.type === "pattern" && (
                  <p className="text-red-600 text-sm">Invalid email format</p>
                )}
              </div>
              <Button
                gradientDuoTone="greenToBlue"
                className="w-full mt-7"
                size="lg"
                type="submit"
              >
                Send link to email
              </Button>
            </form>
            <div
              className={`text-center flex flex-col space-y-2 ${
                isSuccessful ? "" : "hidden"
              }`}
            >
              <img
                src="/success.svg"
                height={32}
                width={32}
                className="mx-auto"
              />
              <div className="text-xl font-medium">Email sent</div>
              <h3 className="mb-5 text-sm font-normal text-gray-500 dark:text-gray-400">
                Check your email and open the link we sent to continue
              </h3>
            </div>
          </div>
        </div>
      </div>
    </section>
  );
};

export default ForgotPassword;
