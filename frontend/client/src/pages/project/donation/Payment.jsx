import { useState, useEffect } from "react";
import { useLocation, useNavigate } from "react-router-dom";
import { loadStripe } from "@stripe/stripe-js";
import { Elements } from "@stripe/react-stripe-js";
import backingService from "../../../services/backing";
import { Outlet } from "react-router-dom";

const stripePromise = loadStripe(import.meta.env.VITE_STRIPE_PK);
console.log(stripePromise);

const Payment = () => {
  const location = useLocation();
  const { amount, user_id, project_id, word_of_support } = location.state;
  const [clientSecret, setClientSecret] = useState("");

  useEffect(() => {
    console.log("amount", amount);
    const fetchIntent = async () => {
      try {
        const response = await backingService.fetchPaymentIntent(
          Number(amount)
        );
        setClientSecret(response.client_secret);
        console.log("client secret", clientSecret);
      } catch (err) {
        console.error(err);
      }
    };

    if (amount) {
      fetchIntent();
    }
  }, []);

  const appearance = {
    theme: "stripe",
  };
  const loader = "auto";

  return (
    <div>
      {clientSecret && (
        <Elements
          options={{ clientSecret, appearance, loader }}
          stripe={stripePromise}
        >
          <Outlet context={{ amount, user_id, project_id, word_of_support }}/>
        </Elements>
      )}
    </div>
  );
};

export default Payment;
