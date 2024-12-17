import { useState } from "react";
import { PiChatsCircle } from "react-icons/pi";
import ragService from "../services/rag";
import { WiStars } from "react-icons/wi";
import { Tooltip } from "flowbite-react";

const Chatbot = () => {
  const [displayChat, setDisplayChat] = useState(false);
  const [chatLog, setChatLog] = useState([]);
  const [input, setInput] = useState();

  const handleChat = async (e) => {
    e.preventDefault();
    setChatLog((prev) => [...prev, { message: input, isUser: true }]);
    setInput("");

    try {
      const response = await ragService.ask(input);
      setChatLog((prev) => [
        ...prev,
        { message: response.answer, isUser: false },
      ]);
    } catch (err) {
      setChatLog((prev) => [
        ...prev,
        {
          message: "There is a problem processing chat request.",
          isUser: false,
        },
      ]);
      console.error(err);
    }
  };

  return (
    <div className="fixed bottom-0 right-0">
      <Tooltip content="AI Assistant">
        <div
          hidden={displayChat}
          onClick={() => setDisplayChat(true)}
          className="rounded-full text-white bg-cyan-800 p-3 mr-5 mb-10 border shadow-md cursor-pointer hover:scale-110"
        >
          <PiChatsCircle className="w-14 h-14" />
        </div>
      </Tooltip>
      <div
        className="bg-white shadow-lg rounded-lg w-[22rem] mr-4"
        hidden={!displayChat}
      >
        <div
          className="flex items-center justify-between py-2 px-4 bg-gradient-to-r from-indigo-600 to-cyan-600 rounded-t-lg cursor-pointer"
          onClick={() => setDisplayChat((val) => !val)}
        >
          <h3 className="text-white flex font-semibold">
            AI Assistant <WiStars className="w-6 h-6" />
          </h3>
          <button className="text-gray-200 hover:text-gray-50 text-xl rounded-full">
            &times;
          </button>
        </div>
        <div className="p-4 h-80 overflow-y-auto border border-gray-300" id="chatbox">
          {chatLog?.map((c) => (
            <ChatMessage message={c.message} isUser={c.isUser} />
          ))}
        </div>
        <form className="p-2 border border-gray-300" onSubmit={handleChat}>
          <input
            type="text"
            placeholder="Ask something..."
            className="w-full p-2 border border-gray-300 rounded-lg"
            onChange={(e) => setInput(e.target.value)}
            value={input}
          />
        </form>
      </div>
    </div>
  );
};

const ChatMessage = ({ message, isUser }) => {
  return (
    <div className={`mb-2 flex justify-${isUser ? "end" : "start"}`}>
      <p className={`bg-${isUser ? "gray-200" : "cyan-100"} px-3 py-2 rounded-lg`}>
        {message}
      </p>
    </div>
  );
};

export default Chatbot;
