import { Avatar, Button, FileInput, Modal } from "flowbite-react";
import { useEffect, useState } from "react";
import { useForm } from "react-hook-form";
import { useAuthContext } from "../../context/AuthContext";
import userService from "../../services/user";
import uploadService from "../../services/upload";
import { useNavigate } from "react-router-dom";
import utils from "../../utils/utils";

const Settings = () => {
  const { token, user } = useAuthContext();
  const navigate = useNavigate();
  const [img, setImg] = useState();
  const [preview, setPreview] = useState();
  const [isSuccessful, setIsSuccessful] = useState({
    status: false,
    field: null,
  });
  const [isFailed, setIsFailed] = useState({ status: false, field: null });
  //const [failReason, setFailReason] = useState();

  const {
    register: registerName,
    handleSubmit: handleSubmitName,
    formState: { errors: errorsName },
    reset: resetName,
  } = useForm();

  const {
    register: registerEmail,
    handleSubmit: handleSubmitEmail,
    formState: { errors: errorsEmail },
    reset: resetEmail,
  } = useForm();

  const {
    register: registerPassword,
    handleSubmit: handleSubmitPassword,
    formState: { errors: errorsPassword },
    getValues: getValuesPassword,
    reset: resetPassword,
  } = useForm();

  const [openEditName, setOpenEditName] = useState(false);
  const [openEditEmail, setOpenEditEmail] = useState(false);
  const [openEditPassword, setOpenEditPassword] = useState(false);
  const [openEditAvatar, setOpenEditAvatar] = useState(false);

  const handleUpdateName = async (data) => {
    try {
      await userService.update(token, {
        first_name: data.first_name,
        last_name: data.last_name,
      });
      // notify success
      setIsSuccessful({ status: true, field: "name" });
      setTimeout(() => {
        navigate(0);
      }, 1000);
    } catch (err) {
      setIsFailed({ status: true, field: "name" });
      console.error(err);
    }
  };

  const handleUpdateEmail = async (data) => {
    try {
      await userService.update(token, {
        email: data.email,
      });
      // notify success
      setIsSuccessful({ status: true, field: "email" });
      setTimeout(() => {
        navigate(0);
      }, 1000);
    } catch (err) {
      setIsFailed({ status: true, field: "email" });
      console.error(err);
    }
  };

  const handleUpdatePassword = async (data) => {
    try {
      await userService.updatePassword(token, {
        old_password: data.old_password,
        new_password: data.new_password,
      });
      // notify success
      setIsSuccessful({ status: true, field: "password" });
      setTimeout(() => {
        navigate(0);
      }, 1000);
    } catch (err) {
      setIsFailed({ status: true, field: "password" });
      console.error(err);
    }
  };

  const handleUpdateAvatar = async () => {
    try {
      const imageUrl = await utils.uploadImage(img);
      await userService.update(token, {
        profile_picture: imageUrl,
      });
      // notify success
      setIsSuccessful({ status: true, field: "avatar" });
      setTimeout(() => {
        navigate(0);
      }, 1000);
    } catch (err) {
      setIsFailed({ status: true, field: "avatar" });
      console.error(err);
    }
  };

  useEffect(() => {
    if (!img || img?.length == 0) {
      setPreview(undefined);
      return;
    }

    const objectUrl = URL.createObjectURL(img[0]);
    setPreview(objectUrl);

    return () => URL.revokeObjectURL(img[0]);
  }, [img]);

  function onSelectImage(e) {
    if (!e.target.files || e.target.files.length == 0) {
      setImg(undefined);
      return;
    }

    setImg(e.target.files);
  }

  return (
    <section className="py-10 flex flex-col items-center">
      <div className="text-3xl font-semibold mb-5">Account Settings</div>
      <div className="w-1/3 px-10 py-5 space-y-4">
        <div className="grid grid-cols-12 py-3">
          <div className="col-span-10 space-y-2 ">
            <div className="font-medium">Name</div>
            <div>{`${user?.first_name} ${user?.last_name}`}</div>
          </div>
          <Button
            color="light"
            className="col-span-2 h-fit"
            size="md"
            onClick={() => setOpenEditName(true)}
          >
            Edit
          </Button>
        </div>
        <Modal
          show={openEditName}
          size="md"
          onClose={() => {
            setOpenEditName(false);
            resetName();
          }}
          popup
        >
          <Modal.Header />
          <Modal.Body>
            <form
              onSubmit={handleSubmitName(handleUpdateName)}
              className="space-y-4"
            >
              <div className="space-y-2">
                <div>
                  <label className="block mb-2 font-medium text-gray-900 text-sm">
                    First Name:
                  </label>
                  <input
                    {...registerName("first_name", {
                      required: "First name is required",
                    })}
                    type="text"
                    className="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5"
                    placeholder={user?.first_name}
                  />
                  {errorsName.first_name?.type === "required" && (
                    <p className="text-red-600 text-sm">
                      {errorsName.first_name.message}
                    </p>
                  )}
                </div>
                <div>
                  <label className="block mb-2 font-medium text-gray-900 text-sm">
                    Last Name:
                  </label>
                  <input
                    {...registerName("last_name", {
                      required: "Last name is required",
                    })}
                    type="text"
                    className="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5"
                    placeholder={user?.last_name ? user.last_name : ""}
                  />
                  {errorsName.last_name?.type === "required" && (
                    <p className="text-red-600 text-sm">
                      {errorsName.last_name.message}
                    </p>
                  )}
                </div>
              </div>
              <Button type="submit" color="dark" className="w-full">
                Change name
              </Button>
            </form>
          </Modal.Body>
        </Modal>

        <hr />

        <div className="grid grid-cols-12 py-3">
          <div className="col-span-10 space-y-2 ">
            <div className="font-medium">Email</div>
            <div>{user?.email}</div>
          </div>
          <Button
            color="light"
            className="col-span-2 h-fit"
            size="md"
            onClick={() => setOpenEditEmail(true)}
          >
            Edit
          </Button>
        </div>
        <Modal
          show={openEditEmail}
          size="md"
          onClose={() => {
            setOpenEditEmail(false);
            resetEmail();
          }}
          popup
        >
          <Modal.Header />
          <Modal.Body>
            <form
              onSubmit={handleSubmitEmail(handleUpdateEmail)}
              className="space-y-4"
            >
              <div>
                <label className="block mb-2 font-medium text-gray-900 text-sm">
                  Email Address:
                </label>
                <input
                  {...registerEmail("email", {
                    required: "Email is required",
                  })}
                  type="email"
                  className="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5"
                  placeholder=""
                />
                {errorsName.email?.type === "required" && (
                  <p className="text-red-600 text-sm">
                    {errorsEmail.email.message}
                  </p>
                )}
              </div>
              <Button type="submit" color="dark" className="w-full">
                Change email
              </Button>
            </form>
          </Modal.Body>
        </Modal>

        <hr />

        <div className="grid grid-cols-12 py-3">
          <div className="col-span-10 space-y-2 ">
            <div className="font-medium">Password</div>
            <div className="">••••••••</div>
          </div>
          <Button
            color="light"
            className="col-span-2 h-fit"
            size="md"
            onClick={() => setOpenEditPassword(true)}
          >
            Edit
          </Button>
        </div>
        <Modal
          show={openEditPassword}
          size="md"
          onClose={() => {
            setOpenEditPassword(false);
            resetPassword();
          }}
          popup
        >
          <Modal.Header />
          <Modal.Body>
            <form
              onSubmit={handleSubmitPassword(handleUpdatePassword)}
              className="space-y-4"
            >
              <div className="space-y-2">
                <div>
                  <label className="block mb-2 font-medium text-gray-900 text-sm">
                    Old Password:
                  </label>
                  <input
                    {...registerPassword("old_password", {
                      required: "First name is required",
                    })}
                    type="password"
                    className="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5"
                    placeholder=""
                  />
                  {errorsPassword.old_password?.type === "required" && (
                    <p className="text-red-600 text-sm">
                      {errorsPassword.old_password.message}
                    </p>
                  )}
                </div>
                <div>
                  <label className="block mb-2 font-medium text-gray-900 text-sm">
                    New Password:
                  </label>
                  <input
                    {...registerPassword("new_password", {
                      required: "Enter your new password",
                      minLength: 8,
                      maxLength: 50,
                      validate: (value) => {
                        const { old_password } = getValuesPassword();
                        return (
                          value !== old_password ||
                          "Must be different from old password"
                        );
                      },
                    })}
                    type="password"
                    className="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5"
                    placeholder=""
                  />
                  {errorsPassword.new_password?.type === "required" && (
                    <p className="text-red-600 text-sm">
                      {errorsPassword.new_password.message}
                    </p>
                  )}
                  {errorsPassword.new_password?.type === "minLength" && (
                    <p className="text-red-600 text-sm">Minimum length is 8</p>
                  )}
                  {errorsPassword.new_password?.type === "maxLength" && (
                    <p className="text-red-600 text-sm">Maximum length is 50</p>
                  )}
                  {errorsPassword.new_password?.type === "validate" && (
                    <p className="text-red-600 text-sm">
                      {errorsPassword.new_password.message}
                    </p>
                  )}
                </div>
                <div>
                  <label className="block mb-2 font-medium text-gray-900 text-sm">
                    Confirm New Password:
                  </label>
                  <input
                    {...registerPassword("re_new_password", {
                      required: "Confirm your password",
                      validate: (value) => {
                        const { new_password } = getValuesPassword();
                        return (
                          value === new_password || "Password does not match"
                        );
                      },
                    })}
                    type="password"
                    className="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5"
                    placeholder=""
                  />
                  {errorsPassword.re_new_password?.type === "required" && (
                    <p className="text-red-600 text-sm">
                      {errorsPassword.re_new_password.message}
                    </p>
                  )}
                  {errorsPassword.re_new_password?.type === "validate" && (
                    <p className="text-red-600 text-sm">
                      {errorsPassword.re_new_password.message}
                    </p>
                  )}
                </div>
              </div>
              <Button type="submit" color="dark" className="w-full">
                Change password
              </Button>
            </form>
          </Modal.Body>
        </Modal>

        <hr />

        <div className="py-3">
          <div className="space-y-2">
            <div className="font-medium">Profile picture</div>
            {user?.profile_picture && !openEditAvatar && (
              <div className="grid grid-cols-12 py-3">
                <div className="col-span-10 overflow-hidden space-y-4 flex flex-col items-start">
                  <Avatar img={user.profile_picture} size="xl" rounded />
                </div>
                <Button
                  color="light"
                  className="col-span-2 row-span-1 h-fit"
                  size="md"
                  onClick={() => setOpenEditAvatar(true)}
                >
                  Edit
                </Button>
              </div>
            )}
            {(!user?.profile_picture || openEditAvatar) && (
              <div className="flex flex-col items-start space-y-4 py-3">
                {img && (
                  <div className="grid grid-cols-10">
                    <div className="overflow-hidden col-span-3">
                      <Avatar img={preview} size="xl" rounded />
                    </div>
                    <div className="col-span-7 flex items-baseline space-x-1">
                      <Button color="light" onClick={handleUpdateAvatar}>
                        Confirm upload
                      </Button>
                      {user?.profile_picture && (
                        <Button
                          color="light"
                          onClick={() => setOpenEditAvatar(false)}
                        >
                          Cancel
                        </Button>
                      )}
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

          {/*<div className="my-3 space-y-3">
            <p className="text-sm text-gray-600">
              Delete your account will remove all your fundraising projects and
              you will not longer be able to sign in with this account
            </p>
            <Button color="red" size="lg" className="w-full">
              Delete account
            </Button>
          </div>
          */}
        </div>
      </div>
      <Modal
        show={isSuccessful.status}
        size="md"
        onClose={() => setIsSuccessful({ status: false, field: null })}
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
              Successfully updated your {isSuccessful.field}
            </h3>
          </div>
        </Modal.Body>
      </Modal>
      <Modal
        show={isFailed.status}
        size="md"
        onClose={() => setIsFailed({ status: false, field: null })}
        popup
      >
        <Modal.Header />
        <Modal.Body>
          <div className="text-center flex flex-col space-y-2">
            <img src="/fail.svg" height={32} width={32} className="mx-auto" />
            <h3 className="mb-5 text-lg font-normal text-gray-500 dark:text-gray-400">
              Failed to update your {isFailed.field}
            </h3>
          </div>
        </Modal.Body>
      </Modal>
    </section>
  );
};

export default Settings;
