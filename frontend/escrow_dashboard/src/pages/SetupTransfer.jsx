import { Button } from "@tremor/react";
import { Label } from "flowbite-react";
import { useForm } from "react-hook-form";
import Cleave from "cleave.js/react";
import utils from "../utils/utils";
import cardService from "../services/card";
import { useAuthContext } from "../context/AuthContext";
import { useNavigate } from "react-router-dom";

const SetupTransfer = () => {
  const { token } = useAuthContext();
  const navigate = useNavigate();
  const {
    register,
    handleSubmit,
    setValue,
    setError,
    formState: { errors },
  } = useForm();
  const onSubmit = (data) => {
    console.log(data);
    const setupTransfer = async (data) => {
      try {
        const payload = {
          card_number: data.card_number,
          expiration_month: utils.parseExpiry(data.expiry).month,
          expiration_year: utils.parseExpiry(data.expiry).year,
          cvv: data.cvv,
          brand: data.brand,
          owner_name: data.owner_name,
        };

        await cardService.setupTransfer(token, payload);
        navigate("/");
      } catch (err) {
        console.error(err);
        setError("server", {
          type: "manual",
          message: "Invalid credit card information",
        });
      }
    };

    setupTransfer(data);
  };

  return (
    <section className="bg-slate-200 flex flex-col items-center justify-center min-h-screen w-full space-y-10">
      <div className="rounded-lg border py-10 px-14 w-1/3 bg-white">
        <div className="flex flex-col items-center mb-10 space-y-1">
          <div className="text-3xl font-semibold">Welcome to Donorbox!</div>
          <div className="text-lg text-gray-700">
            Setup your transfer details to start accepting funds
          </div>
        </div>
        <form className="border-t pt-5" onSubmit={handleSubmit(onSubmit)}>
          {errors.server?.type === "manual" && (
            <p className="text-red-600 text-sm">{errors.server.message}</p>
          )}
          <div className="space-y-4">
            <div className="space-y-1">
              <Label className="text-gray-700">Card brand:</Label>
              <div className="flex space-x-3">
                <div className="flex items-center py-1 px-5 border border-gray-300 rounded-lg dark:border-gray-700">
                  <input
                    {...register("brand")}
                    checked
                    id="visa"
                    type="radio"
                    value="VISA"
                    name="card-radio"
                    className="w-4 h-4 text-blue-600 bg-gray-100 border-gray-300 focus:ring-blue-500 dark:focus:ring-blue-600 dark:ring-offset-gray-800 focus:ring-2 dark:bg-gray-700 dark:border-gray-600"
                  />
                  <label
                    htmlFor="visa"
                    className="w-full ms-2 text-sm font-medium text-gray-900 dark:text-gray-300"
                  >
                    <img src="/visa.svg" className="h-14" />
                  </label>
                </div>
                <div className="flex items-center py-1 px-5 border border-gray-300 rounded-lg dark:border-gray-700">
                  <input
                    {...register("brand")}
                    id="mastercard"
                    type="radio"
                    value="MASTERCARD"
                    name="card-radio"
                    className="w-4 h-4 text-blue-600 bg-gray-100 border-gray-300 focus:ring-blue-500 dark:focus:ring-blue-600 dark:ring-offset-gray-800 focus:ring-2 dark:bg-gray-700 dark:border-gray-600"
                  />
                  <label
                    htmlFor="mastercard"
                    className="w-full ms-2 text-sm font-medium text-gray-900 dark:text-gray-300"
                  >
                    <img src="/mastercard.svg" className="h-14" />
                  </label>
                </div>
              </div>
            </div>
            <div className="space-y-1">
              <Label>Card number:</Label>
              <Cleave
                className="rounded-lg border-gray-300 w-full"
                options={{ creditCard: true }}
                {...register("card_number", {
                  required: "Card number is required",
                  minLength: 13,
                  maxLength: 16,
                })}
                onChange={(e) =>
                  setValue("card_number", e.target.value.replace(/\s/g, ""))
                }
              />
              {errors.card_number?.type === "required" && (
                <p className="text-red-600 text-sm">
                  errors.card_number.message
                </p>
              )}
              {errors.card_number?.type === "minLength" && (
                <p className="text-red-600 text-sm">Invalid card number</p>
              )}
            </div>
            <div className="grid grid-cols-2 gap-4">
              <div className="space-y-1">
                <Label>Expiration date:</Label>
                <Cleave
                  className="rounded-lg border-gray-300 w-full"
                  placeholder="MM/YY"
                  options={{ date: true, datePattern: ["m", "y"] }}
                  {...register("expiry", {
                    required: "Expiration date is required",
                    pattern: /^(0[1-9]|1[0-2])\/?([0-9]{2})$/,
                  })}
                  onChange={(e) =>
                    setValue("expiry", e.target.value.replace("/", ""))
                  }
                />
                {errors.expiry?.type === "required" && (
                  <p className="text-red-600 text-sm">errors.expiry.message</p>
                )}
                {errors.expiry?.type === "pattern" && (
                  <p className="text-red-600 text-sm">Invalid date</p>
                )}
              </div>
              <div className="space-y-1">
                <Label>CVV</Label>
                <Cleave
                  className="rounded-lg border-gray-300 w-full"
                  options={{ blocks: [3], numericOnly: true }}
                  {...register("cvv", {
                    required: "CVV is required",
                    minLength: 3,
                    maxLength: 3,
                  })}
                  onChange={(e) => setValue("cvv", e.target.value)}
                />
                {errors.cvv?.type === "required" && (
                  <p className="text-red-600 text-sm">{errors.cvv.message}</p>
                )}
                {errors.cvv?.type === "minLength" && (
                  <p className="text-red-600 text-sm">Invalid CVV</p>
                )}
              </div>
            </div>
            <div className="space-y-1">
              <Label>Card owner name:</Label>
              <input
                className="rounded-lg border-gray-300 w-full ring-blue-700"
                {...register("owner_name", {
                  required: "Name is required",
                })}
                placeholder="JOHN DOE"
              />
              {errors.owner_name?.type === "required" && (
                <p className="text-red-600 text-sm">
                  {errors.owner_name.message}
                </p>
              )}
            </div>

            <div className="flex justify-center">
              <Button type="submit" size="xl" className="w-1/3">
                Submit
              </Button>
            </div>
          </div>
        </form>
      </div>
    </section>
  );
};

export default SetupTransfer;
