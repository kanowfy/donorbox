import { Button, Checkbox, Textarea } from "flowbite-react";
import { useState } from "react";

const Donate = () => {
  const [wosChecked, setWosChecked] = useState(false);

  function handleSelectWOS() {
    setWosChecked(!wosChecked);
  }
  return (
    <section className="py-10 flex flex-col items-center">
      <div>
        <a
          href="#"
          className="flex items-center mb-6 text-2xl font-semibold text-gray-800 tracking-tight"
        >
          <img className="w-8 h-8 mr-2" src="/logo.svg" alt="logo" />
          Donorbox
        </a>
      </div>
      <div className="w-2/5 rounded-lg shadow-lg px-10 py-5 space-y-4">
        <div className="grid grid-cols-12 space-x-2">
          <div className="col-span-3">
            <img
              src="https://images.pexels.com/photos/2253879/pexels-photo-2253879.jpeg"
              className="h-28 aspect-[4/3]"
            />
          </div>
          <div className="col-span-9">
            <span>You are supporting </span>
            <span className="font-semibold">
              Lorem ipsum dolor sit amet consectetur adipisicing elit.
            </span>
          </div>
        </div>
        <div>
          <label className="block mb-2 font-medium text-gray-900">
            Select amount to donate:
          </label>
          <div className="bg-gray-50 border border-gray-300 text-gray-900 rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full px-5 py-3">
            <span>₫</span>
            <input
              type="text"
              className="border-0 focus:ring-0 focus:border-0 bg-gray-50 w-11/12"
              placeholder=""
            />
          </div>
        </div>

        <div className="font-semibold">Enter payment information:</div>
        <div className="border rounded-lg p-3 space-y-3">
          <div>
            <label className="block mb-2 text-sm font-medium text-gray-900">
              Card number
            </label>
            <input
              type="email"
              className="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full py-3 px-4"
              placeholder=""
            />
          </div>
          <div className="flex gap-4">
            <div className="w-1/2">
              <label className="block mb-2 text-sm font-medium text-gray-900">
                Expiration date
              </label>
              <input
                type="text"
                className="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full py-3 px-4"
                placeholder="MM/YY"
              />
            </div>
            <div className="w-1/2">
              <label className="block mb-2 text-sm font-medium text-gray-900">
                CVV
              </label>
              <input
                type="text"
                className="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full py-3 px-4"
                placeholder="XXX"
              />
            </div>
          </div>
        </div>

        <div className="space-x-2 ml-2">
          <Checkbox onChange={handleSelectWOS} checked={wosChecked} />
          <label className="font-medium text-gray-700">
            Write word of support
          </label>
        </div>

        {wosChecked && (
          <div>
            <Textarea placeholder="Enter word of support" rows={2} />
          </div>
        )}

        <Button color="success" className="w-full" size={"xl"} type="submit">
          Donate
        </Button>
      </div>
    </section>
  );
};

export default Donate;
