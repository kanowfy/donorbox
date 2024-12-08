import { useEffect, useState } from "react";
import { useAuthContext } from "../context/AuthContext";
import notificationService from "../services/notification";
import { BASE_URL } from "../constants";
import { Dropdown } from "flowbite-react";
import { IoMdNotificationsOutline } from "react-icons/io";
import { FaBell } from "react-icons/fa";
import { Link } from "react-router-dom";
import utils from "../utils/utils";

const Notification = () => {
  const { user, token } = useAuthContext();
  const [notifs, setNotifs] = useState();
  const [notifsLen, setNotifsLen] = useState(0);

  useEffect(() => {
    const fetchNotifs = async () => {
      try {
        const response = await notificationService.getNotifications(
          token,
          user?.id
        );
        setNotifs(response.notifications);
        setNotifsLen(response.notifications?.filter(n => !n.is_read).length || 0);
        console.log(response.notifications);
      } catch (err) {
        console.error(err);
      }
    };

    fetchNotifs();
  }, []);

  useEffect(() => {
    const eventSource = new EventSource(`${BASE_URL}/notifications/events`);
    eventSource.onmessage = function (event) {
      const notif = JSON.parse(event.data);
      setNotifs((prev) => [...prev, notif]);
      setNotifsLen(prev => prev+1);
    };

    return () => {
      eventSource.close();
    };
  }, []);

  const handleReadNotif = async (notif) => {
    try {
      await notificationService.updateReadNotification(token, notif.id);
    } catch (err) {
      console.error(err);
    }
  };

  return (
    <Dropdown
      theme={{
        floating: {
            item: {
                container: "",
                base: "text-sm hover:bg-gray-100 focus:bg-gray-100 text-left px-5 py-2"
            }
        }
      }}
      label=""
      dismissOnClick={true}
      renderTrigger={() => (
        <button className="mr-8 rounded-full mb-3 w-7">
          <div className={`${notifsLen === 0 && "invisible"} absoslute bg-red-500 py-0.5 text-white text-xs rounded-full top-0 right-0 transform translate-x-1/2 translate-y-1/2`}>
            {notifsLen}
          </div>
          <div>
            <IoMdNotificationsOutline className="h-full w-full" />
          </div>
        </button>
      )}
    >
      { notifsLen > 0 ? (
      notifs?.map((n) => (
        <Link to={`/fundraiser/${n.project_id}`} key={n.id}>
          <Dropdown.Item
            className={`max-w-sm ${
              n.is_read ? "text-gray-500" : "text-gray-900"
            }`}
            onClick={() => handleReadNotif(n)}
          >
            <div className="flex justify-between">
            <div>
                {n.type === "approved_project" || n.type === "rejected_project" ? (
                    <div className="font-semibold text-blue-800">Campaign Submission Result</div>
                ) : (
                    <div className="font-semibold text-blue-800">Milestone Resolved</div>
                )}
            </div>
                <div className="text-xs text-gray-500">
                  {utils.formatDate(
                    new Date(utils.parseDateFromRFC3339(n.created_at))
                  )}
                </div>
            </div>
            <div className="mt-2">
                {n.message}
            </div>

          </Dropdown.Item>
        </Link>
      ))) : (
        <div className="px-5 py-4 flex">
          <div><FaBell className="w-5 h-5 mr-1 mt-1"/></div>
        <div className="text-gray-700">You don't have any notifications</div>
        </div>
      )
      }
    </Dropdown>
  );
};

export default Notification;
