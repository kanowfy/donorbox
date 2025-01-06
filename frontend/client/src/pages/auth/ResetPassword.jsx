import { Button } from "flowbite-react";
import { useForm } from "react-hook-form";
import user from "../../services/user";
import { useState } from "react";
import { useNavigate, useSearchParams } from "react-router-dom";

const ResetPassword = () => {
  const [isSuccessful, setIsSuccessful] = useState(false);
  const [params] = useSearchParams();
  const navigate = useNavigate();
  const {
    register,
    formState: { errors },
    handleSubmit,
    getValues,
  } = useForm();

  const onSubmit = async (data) => {
    try {
      await user.resetPassword(params.get("token"), data.password);
      setIsSuccessful(true);
      setTimeout(() => {
        navigate("/login");
      }, 1500);
    } catch (err) {
      console.error(err);
    }
  };

  return (
    <section className="bg-gradient-to-tr from-sky-200 to-sky-400">
      <div className="flex flex-col items-center justify-center px-6 py-8 mx-auto md:h-screen lg:py-0">
        <div className="w-full bg-white rounded-lg shadow md:mt-0 sm:max-w-md xl:p-0">
          <div className="p-6 space-y-4 md:space-y-6 sm:p-8">
            <form onSubmit={handleSubmit(onSubmit)} hidden={isSuccessful}>
              <div className="mb-5 font-medium text-lg tracking-tight leading-tight">
                Reset your password
              </div>
              <div>
                <label className="block mb-2 text-sm font-medium text-gray-900">
                  New password <span className="text-red-700">*</span>
                </label>
                <input
                  {...register("new_password", {
                    required: true,
                    minLength: 8,
                    maxLength: 50,
                  })}
                  type="password"
                  placeholder=""
                  className="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5"
                />
                {errors.password?.type === "required" && (
                  <p className="text-red-600 text-sm">Password is required</p>
                )}
                {errors.password?.type === "minLength" && (
                  <p className="text-red-600 text-sm">Minimum length is 8</p>
                )}
                {errors.lastName?.type === "maxLength" && (
                  <p className="text-red-600 text-sm">Maximum length is 50</p>
                )}
              </div>
              <div>
                <label className="block mb-2 text-sm font-medium text-gray-900">
                  Confirm new password <span className="text-red-700">*</span>
                </label>
                <input
                  {...register("password", {
                    required: true,
                    validate: (value) => {
                      const { password } = getValues();
                      return password === value || "Password must match";
                    },
                  })}
                  type="password"
                  placeholder=""
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
                className="w-full mt-7"
                size="lg"
                type="submit"
              >
                Reset Password
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
              <div className="text-xl font-medium">
                Password reset successfully!
              </div>
              <h3 className="mb-5 text-sm font-normal text-gray-500 dark:text-gray-400">
                You will be redirected to login page...
              </h3>
            </div>
          </div>
        </div>
      </div>
    </section>
  );
};

export default ResetPassword;
