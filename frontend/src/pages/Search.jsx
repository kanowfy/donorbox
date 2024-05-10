import { useState } from "react";
import SearchResults from "../components/SearchResults";

const Search = () => {
  const [query, setQuery] = useState("");

  function submitSearchQuery(e) {
    e.preventDefault();
  }

  return (
    <>
      <section className="min-h-screen bg-cover bg-gradient-to-b from-white to-emerald-100 mt-20">
        <div className="flex flex-col items-center">
          <div className="pb-7 font-semibold text-4xl tracking-tight">
            Search for fundraising projects
          </div>
          <div className="pb-10 font-normal text-lg">
            Find fundraisers by category, name or location
          </div>
          <div>
            <form
              onSubmit={submitSearchQuery}
              className="min-w-128 max-w-md mx-auto"
            >
              <label className="mb-2 text-sm font-medium text-gray-900 sr-only dark:text-white">
                Search
              </label>
              <div className="relative">
                <div className="absolute inset-y-0 start-0 flex items-center ps-3 pointer-events-none">
                  <svg
                    className="w-4 h-4 text-gray-500 dark:text-gray-400"
                    aria-hidden="true"
                    xmlns="http://www.w3.org/2000/svg"
                    fill="none"
                    viewBox="0 0 20 20"
                  >
                    <path
                      stroke="currentColor"
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      strokeWidth="2"
                      d="m19 19-4-4m0-7A7 7 0 1 1 1 8a7 7 0 0 1 14 0Z"
                    />
                  </svg>
                </div>
                <input
                  type="search"
                  className="block w-full p-4 ps-10 text-sm text-gray-900 border border-gray-300 rounded-lg bg-gray-50 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                  onChange={(e) => setQuery(e.target.value)}
                  value={query}
                  placeholder="Search..."
                  required
                />
                <button
                  type="submit"
                  className="text-white absolute end-2.5 bottom-2.5 bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-4 py-2 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
                >
                  Search
                </button>
              </div>
            </form>
          </div>
        </div>
        <section className="mx-32 my-10">
          <SearchResults />
        </section>
      </section>
    </>
  );
};

export default Search;
