import { Button, Label, Modal, TextInput } from "flowbite-react";
import { useEffect, useState } from "react";
import { useOutletContext } from "react-router-dom";
import transferService from "../../../services/transfer";
import { useAuthContext } from "../../../context/AuthContext";
import { useForm } from "react-hook-form";
import Cleave from "cleave.js/react";
import utils from "../../../utils/utils";

/*
const cardExample = {
  card_owner_name: "JOHNNY ENGLISH",
  last_four_digits: "2222",
  brand: "VISA",
};
 */

const ManageTransfer = () => {
  const { token } = useAuthContext();
  const { project } = useOutletContext();
  const [card, setCard] = useState();
  const [setupSuccessful, setSetupSuccessful] = useState(false);
  const [setupFailed, setSetupFailed] = useState(false);
  const [isLoading, setIsLoading] = useState(false);
  const {
    register,
    setValue,
    handleSubmit,
    formState: { errors },
  } = useForm();

  useEffect(() => {
    const getCard = async () => {
      try {
        const response = await transferService.getCard(token, project.card_id);
        setCard(response.card);
      } catch (err) {
        console.error(err);
      }
    };

    if (project?.card_id) {
      getCard();
    }
  }, [project?.card_id, token]);

  const onSubmit = (data) => {
    const setupTransfer = async (data) => {
      setIsLoading(true);
      try {
        const payload = {
          card_number: data.card_number,
          expiration_month: utils.parseExpiry(data.expiry).month,
          expiration_year: utils.parseExpiry(data.expiry).year,
          cvv: data.cvv,
          brand: data.brand,
          owner_name: data.owner_name,
        };

        await transferService.setupCard(token, project.id, payload);
        setIsLoading(false);
        setSetupSuccessful(true);
      } catch (err) {
        console.error(err);
        setSetupFailed(true);
      }
    };

    setupTransfer(data);
  };

  return (
    <div>
      <div>
        <div className="text-3xl font-semibold">Transfer</div>
        <div className="text-gray-600 text-sm mt-1">
          Set up your bank account to receive donations.
        </div>
      </div>
      <div className="my-10">
        {card ? (
          <div className="flex justify-between w-2/3">
            <div className="mb-2 font-medium ">Saved transfer:</div>
            {card?.brand == "VISA" ? (
              <div className="rounded-lg px-6 py-4 border border-black w-fit bg-blue-900 space-y-2">
                <div className="flex justify-end">
                  <img src="/visa.svg" className="h-10" />
                </div>

                <div className="text-white text-lg font-mono tracking-widest">{`•••• •••• •••• ${card?.last_four_digits}`}</div>
                <div className="text-white font-medium font-mono">
                  {card.card_owner_name}
                </div>
              </div>
            ) : (
              <div className="rounded-lg px-6 py-4 border border-black w-fit bg-gray-200 space-y-2">
                <div className="flex justify-end">
                  <img src="/mastercard.svg" className="h-10" />
                </div>

                <div className="text-lg font-mono tracking-widest">{`•••• •••• •••• ${card?.last_four_digits}`}</div>
                <div className="font-medium font-mono">
                  {card.card_owner_name}
                </div>
              </div>
            )}
          </div>
        ) : (
          <div className="w-full space-y-4">
            <div className="font-medium">
              Fill out the form below to complete transfer setup for your
              fundraiser:
            </div>
            <form
              className="border rounded-lg px-7 py-5 w-1/2 shadow-md"
              onSubmit={handleSubmit(onSubmit)}
            >
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
                      <p className="text-red-600 text-sm">
                        errors.expiry.message
                      </p>
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
                      <p className="text-red-600 text-sm">
                        {errors.cvv.message}
                      </p>
                    )}
                    {errors.cvv?.type === "minLength" && (
                      <p className="text-red-600 text-sm">Invalid CVV</p>
                    )}
                  </div>
                </div>
                <div className="space-y-1">
                  <Label>Card owner name:</Label>
                  <TextInput
                    theme={{
                      field: {
                        input: {
                          base: "block w-full border disabled:cursor-not-allowed disabled:opacity-50 border-gray-300",
                        },
                      },
                    }}
                    color="blue"
                    className="w-full"
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
                  <Button
                    type="submit"
                    color="blue"
                    className="w-1/3"
                    isProcessing={isLoading}
                  >
                    Submit
                  </Button>
                </div>
              </div>
            </form>
            <Modal
              show={setupSuccessful}
              size="md"
              onClose={() => setSetupSuccessful(false)}
              popup
            >
              <Modal.Header />
              <Modal.Body>
                <div className="text-center flex flex-col space-y-2">
                  <img
                    src="/success.svg"
                    height={32}
                    width={32}
                    className="mx-auto"
                  />
                  <h3 className="mb-5 text-lg font-normal text-gray-500 dark:text-gray-400">
                    Transfer setup successful! You can receive donations once
                    your fundraiser succeeds
                  </h3>
                </div>
              </Modal.Body>
            </Modal>
            <Modal
              show={setupFailed}
              size="md"
              onClose={() => setSetupFailed(false)}
              popup
            >
              <Modal.Header />
              <Modal.Body>
                <div className="text-center flex flex-col space-y-2">
                  <img
                    src="/fail.svg"
                    height={32}
                    width={32}
                    className="mx-auto"
                  />
                  <h3 className="mb-5 text-lg font-normal text-gray-500 dark:text-gray-400">
                    Failed to setup transfer. Make sure you have entered correct
                    transfer details.
                  </h3>
                </div>
              </Modal.Body>
            </Modal>
          </div>
        )}
      </div>
    </div>
  );
};

export default ManageTransfer;
