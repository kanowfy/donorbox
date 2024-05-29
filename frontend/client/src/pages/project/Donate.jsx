import {
  Button,
  Checkbox,
  Label,
  Modal,
  Textarea,
  TextInput,
  Tooltip,
} from "flowbite-react";
import { useEffect, useState } from "react";
import { useForm } from "react-hook-form";
import Cleave from "cleave.js/react";
import { useAuthContext } from "../../context/AuthContext";
import utils from "../../utils/utils";
import backingService from "../../services/backing";
import { Link, useNavigate, useParams } from "react-router-dom";
import projectService from "../../services/project";
import { FaAngleLeft } from "react-icons/fa6";
import { BsFillQuestionCircleFill } from "react-icons/bs";

const Donate = () => {
  const params = useParams();
  const navigate = useNavigate();
  const [wosChecked, setWosChecked] = useState(false);
  const { user, loading } = useAuthContext() || {};
  const {
    register,
    handleSubmit,
    setValue,
    formState: { errors },
    reset,
  } = useForm();
  const [isSuccessful, setIsSuccessful] = useState(false);
  const [isFailed, setIsFailed] = useState(false);
  const [project, setProject] = useState();

  const handleSelectWOS = () => {
    setWosChecked(!wosChecked);
  };

  useEffect(() => {
    const fetchProject = async () => {
      try {
        const response = await projectService.getOne(params.id);
        setProject(response.project);
      } catch (err) {
        navigate("/not-found");
        console.error(err);
      }
    };

    if (!loading) {
      fetchProject();
    }
  }, [params.id, user, loading, navigate]);

  const onSubmit = (data) => {
    const donate = async (data) => {
      try {
        const payload = {
          amount: data.amount,
          card_info: {
            card_number: data.card_number,
            expiration_month: utils.parseExpiry(data.expiry).month,
            expiration_year: utils.parseExpiry(data.expiry).year,
            cvv: data.cvv,
            brand: data.brand,
            owner_name: data.owner_name,
          },
        };

        if (data.word_of_support) {
          payload.word_of_support = data.word_of_support;
        }

        if (user) {
          payload.user_id = user.id;
        }

        await backingService.backProject(params.id, payload);
        setIsSuccessful(true);
        setTimeout(() => {
          navigate(`/fundraiser/${params.id}`);
        }, 5000);
      } catch (err) {
        setIsFailed(true);
        console.error(err);
        reset();
      }
    };

    donate(data);
  };

  return (
    <section className="py-10 flex flex-col items-center bg-gray-100 min-h-screen">
      <div className="w-full grid grid-cols-3">
        <Link
          to={`/fundraiser/${params?.id}`}
          className="flex hover:underline justify-center"
        >
          <FaAngleLeft className="w-7 h-7" />
          <div>Back to Fundraiser</div>
        </Link>
        <Link
          to={`/fundraisers/${params?.id}`}
          className="flex items-center justify-center mb-6 text-2xl font-semibold text-gray-800 tracking-tight"
        >
          <img className="w-8 h-8 mr-2" src="/logo.svg" alt="logo" />
          Donorbox
        </Link>
        <div></div>
      </div>
      <form
        className="w-1/3 rounded-lg shadow-lg px-10 py-5 space-y-4 bg-white"
        onSubmit={handleSubmit(onSubmit)}
      >
        <div className="grid grid-cols-12 space-x-2">
          <div className="col-span-3 rounded-xl overflow-hidden h-20 aspect-[4/3] object-cover">
            <img
              src={project?.cover_picture}
              className="w-full h-full m-auto object-cover"
            />
          </div>
          <div className="col-span-9">
            <span>You are supporting </span>
            <span className="font-semibold">{project?.title}</span>
          </div>
        </div>
        <div>
          <label className="block mb-2 font-medium text-gray-900">
            Select amount to donate:
          </label>
          <div className="flex items-baseline bg-gray-50 border border-gray-300 text-gray-900 rounded-lg focus:ring-primary-600 focus:border-primary-600 w-full px-5 py-3">
            <span className="block mb-1 font-medium">₫</span>
            <Cleave
              options={{
                numeral: true,
                numericOnly: true,
                numeralThousandsGroupStyle: "thousand",
                numeralPositiveOnly: true,
              }}
              {...register("amount", {
                required: "Donation amount is required",
                min: 50000,
                max: 100000000,
              })}
              onChange={(e) =>
                setValue("amount", e.target.value.replace(/,/g, ""))
              }
              className="border-0 focus:ring-0 focus:border-0 bg-gray-50 w-11/12 text-xl font-medium autofill:bg-gray-50"
              placeholder=""
            />
          </div>
          {errors.amount?.type === "required" && (
            <p className="text-red-600 text-sm">errors.amount.message</p>
          )}
          {errors.amount?.type === "min" && (
            <p className="text-red-600 text-sm">
              Minimum donation amount is ₫50,000
            </p>
          )}
          {errors.amount?.type === "max" && (
            <p className="text-red-600 text-sm">
              Maximum donation amount is ₫100,000,000
            </p>
          )}
        </div>

        <div className="font-semibold">Enter payment information:</div>
        <div className="rounded-lg border p-3 space-y-2">
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
              className="rounded-lg border-gray-300 w-full text-sm"
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
              <p className="text-red-600 text-sm">errors.card_number.message</p>
            )}
            {errors.card_number?.type === "minLength" && (
              <p className="text-red-600 text-sm">Invalid card number</p>
            )}
          </div>
          <div className="grid grid-cols-2 gap-4">
            <div className="space-y-1">
              <Label>Expiration date:</Label>
              <Cleave
                className="rounded-lg border-gray-300 w-full text-sm"
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
                className="rounded-lg border-gray-300 w-full text-sm"
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
        </div>

        <div className="font-medium text-gray-700 flex space-x-1">
          <Checkbox
            onChange={handleSelectWOS}
            checked={wosChecked}
            className="mt-1 mr-1"
          />
          <div>Write word of support</div>
          <div className="mt-1">
            <Tooltip content="Your word of support will be displayed on the fundraiser page once you have made a successful donation.">
              <BsFillQuestionCircleFill className="w-4 h-4" />
            </Tooltip>
          </div>
        </div>

        {wosChecked && (
          <div>
            <Textarea
              {...register("word_of_support", {
                required: false,
                minLength: 2,
              })}
              placeholder="Enter word of support"
              rows={4}
            />
          </div>
        )}

        <Button color="success" className="w-full" size={"xl"} type="submit">
          Donate
        </Button>
      </form>

      <Modal
        show={isSuccessful}
        size="md"
        onClose={() => setIsSuccessful(false)}
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
              Thanks for supporting! You will be navigated to the fundraiser
              page in a short time...
            </h3>
          </div>
        </Modal.Body>
      </Modal>
      <Modal show={isFailed} size="md" onClose={() => setIsFailed(false)} popup>
        <Modal.Header />
        <Modal.Body>
          <div className="text-center flex flex-col space-y-2">
            <img src="/fail.svg" height={32} width={32} className="mx-auto" />
            <h3 className="mb-5 text-lg font-normal text-gray-500 dark:text-gray-400">
              Failed to process payment. Make sure you have entered correct
              payment details.
            </h3>
          </div>
        </Modal.Body>
      </Modal>
    </section>
  );
};

export default Donate;
