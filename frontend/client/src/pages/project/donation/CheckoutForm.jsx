import { useElements, useStripe } from "@stripe/react-stripe-js";
import { useState } from "react";
import { PaymentElement } from "@stripe/react-stripe-js";
import { Button, Modal, Spinner } from "flowbite-react";
import { useNavigate, useOutletContext, useParams } from "react-router-dom";
import backingService from "../../../services/backing";

const CheckoutForm = () => {
  const params = useParams();
  const navigate = useNavigate();
  const stripe = useStripe();
  const elements = useElements();
  const [isLoading, setIsLoading] = useState(false);
  const [isSuccessful, setIsSuccessful] = useState(false);
  const [isFailed, setIsFailed] = useState({
    status: false,
    message: "",
  });

  const { amount, user_id, project_id, word_of_support } = useOutletContext();

  const handleSubmit = async (e) => {
    e.preventDefault();

    if (!stripe || !elements) {
      // Make sure to disable form submission until Stripe.js has loaded.
      return;
    }

    setIsLoading(true);

    const { error } = await stripe.confirmPayment({
      elements,
      confirmParams: {},
      redirect: "if_required",
    });

    if (error) {
      setIsFailed({
        status: true,
        message: error.message,
      });
    } else {
      console.log(
        "amount ",
        amount,
        "user_id",
        user_id,
        "project_id",
        project_id,
        "wos",
        word_of_support,
      );
      setIsSuccessful(true);
      const response = await backingService.createBacking(project_id, {
        amount: Number(amount),
        user_id: Number(user_id),
        word_of_support: word_of_support,
      });
      console.log(response);
    }
    setIsLoading(false);
    setTimeout(() => {
      navigate("/fundraiser/" + params.id);
    }, 5000);
  };

  const paymentElementOptions = {
    layout: "accordion",
  };

  return (
    <div className="bg-gray-100 min-h-screen">
      <div className="mx-auto max-w-2xl py-10">
        <div className="text-4xl font-bold mb-10 text-center">Checkout</div>
        <div className="">
          <form id="payment-form" onSubmit={handleSubmit}>
            <PaymentElement
              id="payment-element"
              options={paymentElementOptions}
            />
            <Button
              disabled={isLoading || !stripe || !elements}
              id="submit"
              color="blue"
              className="mt-5 mx-auto"
              type="submit"
            >
              {isLoading ? <Spinner /> : "Pay now"}
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
          <Modal
            show={isFailed.status}
            size="md"
            onClose={() => setIsFailed(false)}
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
                  Failed to process payment. Make sure you have entered correct
                  payment details.
                </h3>
              </div>
            </Modal.Body>
          </Modal>
        </div>
      </div>
    </div>
  );
};

export default CheckoutForm;
