import { useState } from "react";
import { PiChatsCircle } from "react-icons/pi";
import ragService from "../services/rag";
import { WiStars } from "react-icons/wi";

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
        setChatLog((prev) => [...prev, { message: response.answer, isUser: false }]);
    } catch(err) {
        setChatLog((prev) => [...prev, { message: "There is a problem processing chat request.", isUser: false }]);
        console.error(err);
    }
  };

  return (
    <div class="fixed bottom-0 right-0">
      <div
        hidden={displayChat}
        onClick={() => setDisplayChat(true)}
        className="rounded-full bg-white p-3 mr-5 mb-10 border shadow-md cursor-pointer"
      >
        <PiChatsCircle className="w-16 h-16" />
      </div>
      <div
        className="bg-white border border-gray-300 shadow-lg rounded-lg w-[22rem] mr-4"
        hidden={!displayChat}
      >
        <div
          class="flex items-center justify-between py-2 px-4 bg-green-600 rounded-t-lg cursor-pointer"
          onClick={() => setDisplayChat((val) => !val)}
        >
          <h3 class="text-white flex">AI Assistant <WiStars className="w-6 h-6"/></h3>
          <button class="text-gray-200 hover:text-gray-50 text-xl rounded-full">
            &times;
          </button>
        </div>
        <div class="p-4 h-80 overflow-y-auto" id="chatbox">
          {chatLog?.map((c) => (
            <ChatMessage message={c.message} isUser={c.isUser} />
          ))}
        </div>
        <form class="p-2 border border-gray-300" onSubmit={handleChat}>
          <input
            type="text"
            placeholder="Ask something..."
            class="w-full p-2 border border-gray-300 rounded-lg"
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
    <div class={`mb-2 flex justify-${isUser ? "end" : "start"}`}>
      <p class={`bg-${isUser ? "gray" : "green"}-200 px-3 py-2 rounded-lg`}>
        {message}
      </p>
    </div>
  );
};

export default Chatbot;
