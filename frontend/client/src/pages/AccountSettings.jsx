import { Avatar, Button, FileInput } from "flowbite-react";
import { useEffect, useState } from "react";

const AccountSettings = () => {
  const [hasAvatar] = useState(false);
  const [img, setImg] = useState();
  const [preview, setPreview] = useState();

  useEffect(() => {
    if (!img) {
      setPreview(undefined);
      return;
    }

    const objectUrl = URL.createObjectURL(img);
    setPreview(objectUrl);

    return () => URL.revokeObjectURL(img);
  }, [img]);

  function onSelectImage(e) {
    if (!e.target.files || e.target.files.length == 0) {
      setImg(undefined);
      return;
    }

    setImg(e.target.files[0]);
  }
  return (
    <section className="py-10 flex flex-col items-center">
      <div className="text-3xl font-semibold mb-5">Account Settings</div>
      <div className="w-1/3 px-10 py-5 space-y-4">
        <div className="grid grid-cols-12 py-3">
          <div className="col-span-10 space-y-2 ">
            <div className="font-medium">Name</div>
            <div>John Doe</div>
          </div>
          <Button color="light" className="col-span-2 h-fit" size="md">
            Edit
          </Button>
        </div>

        <hr />

        <div className="grid grid-cols-12 py-3">
          <div className="col-span-10 space-y-2 ">
            <div className="font-medium">Email</div>
            <div>johndoe123@gmail.com</div>
          </div>
          <Button color="light" className="col-span-2 h-fit" size="md">
            Edit
          </Button>
        </div>

        <hr />

        <div className="grid grid-cols-12 py-3">
          <div className="col-span-10 space-y-2 ">
            <div className="font-medium">Password</div>
            <div className="">••••••••</div>
          </div>
          <Button color="light" className="col-span-2 h-fit" size="md">
            Edit
          </Button>
        </div>

        <hr />

        <div className="py-3">
          <div className="space-y-2">
            <div className="font-medium">Profile picture</div>
            {hasAvatar ? (
              <div className="grid grid-cols-12 py-3">
                <div className="col-span-10 overflow-hidden space-y-4 flex flex-col items-start">
                  <Avatar
                    img="https://images.pexels.com/photos/33537/cat-animal-cat-portrait-mackerel.jpg"
                    size="xl"
                    rounded
                  />
                </div>
                <Button
                  color="light"
                  className="col-span-2 row-span-1 h-fit"
                  size="md"
                >
                  Edit
                </Button>
              </div>
            ) : (
              <div className="flex flex-col items-start space-y-4 py-3">
                {img && (
                  <div className="grid grid-cols-10">
                    <div className="overflow-hidden col-span-3">
                      <Avatar img={preview} size="xl" rounded />
                    </div>
                    <div className="col-span-7 flex items-baseline">
                      <Button color="light" pill>
                        Confirm upload
                      </Button>
                    </div>
                  </div>
                )}
                <FileInput
                  accept="image/png, image/jpeg"
                  color="gray"
                  sizing="lg"
                  onChange={onSelectImage}
                />
              </div>
            )}
          </div>

          <hr />

          <div className="my-3 space-y-3">
            <p className="text-sm text-gray-600">
              Delete your account will remove all your fundraising projects and
              you will not longer be able to sign in with this account
            </p>
            <Button color="red" size="lg" className="w-full">
              Delete account
            </Button>
          </div>
        </div>
      </div>
    </section>
  );
};

export default AccountSettings;
