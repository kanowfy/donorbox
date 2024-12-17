export const BASE_URL = import.meta.env.VITE_BACKEND_URL;
export const SERVE_URL = import.meta.env.VITE_SERVE_URL;
export const AUTH_ACTIONS = {
    LOGIN: "login",
    LOGOUT: "logout",
    LOADED: "loaded"
}

export const CategoryIndexMap = {
    Medical: 1,
    Emergency: 2,
    Education: 3,
    Animals: 4,
    Competition: 5,
    Event: 6,
    Environment: 7,
    Travel: 8,
    Business: 9,
};

export const INFURA_APIKEY = import.meta.env.VITE_INFURA_APIKEY;
export const CONTRACT_ADDRESS = import.meta.env.VITE_CONTRACT_ADDRESS;
export const ABI = [
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "uint256",
				"name": "id",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "uint64",
				"name": "project_id",
				"type": "uint64"
			},
			{
				"indexed": false,
				"internalType": "uint64",
				"name": "milestone_id",
				"type": "uint64"
			},
			{
				"indexed": false,
				"internalType": "string",
				"name": "transfer_image_url",
				"type": "string"
			},
			{
				"indexed": false,
				"internalType": "string",
				"name": "transfer_note",
				"type": "string"
			},
			{
				"indexed": false,
				"internalType": "string",
				"name": "created_at",
				"type": "string"
			}
		],
		"name": "MilestoneReleaseStored",
		"type": "event"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "id",
				"type": "uint256"
			},
			{
				"internalType": "uint64",
				"name": "project_id",
				"type": "uint64"
			},
			{
				"internalType": "uint64",
				"name": "milestone_id",
				"type": "uint64"
			},
			{
				"internalType": "string",
				"name": "transfer_image_url",
				"type": "string"
			},
			{
				"internalType": "string",
				"name": "transfer_note",
				"type": "string"
			},
			{
				"internalType": "string",
				"name": "created_at",
				"type": "string"
			}
		],
		"name": "storeFundRelease",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "id",
				"type": "uint256"
			},
			{
				"internalType": "uint64",
				"name": "project_id",
				"type": "uint64"
			},
			{
				"internalType": "uint64",
				"name": "milestone_id",
				"type": "uint64"
			},
			{
				"internalType": "string",
				"name": "transfer_receipt_url",
				"type": "string"
			},
			{
				"internalType": "string",
				"name": "proof_media_url",
				"type": "string"
			},
			{
				"internalType": "string",
				"name": "created_at",
				"type": "string"
			}
		],
		"name": "storeVerifiedProof",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "uint256",
				"name": "id",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "uint64",
				"name": "project_id",
				"type": "uint64"
			},
			{
				"indexed": false,
				"internalType": "uint64",
				"name": "milestone_id",
				"type": "uint64"
			},
			{
				"indexed": false,
				"internalType": "string",
				"name": "transfer_receipt_url",
				"type": "string"
			},
			{
				"indexed": false,
				"internalType": "string",
				"name": "proof_media_url",
				"type": "string"
			},
			{
				"indexed": false,
				"internalType": "string",
				"name": "created_at",
				"type": "string"
			}
		],
		"name": "VerifiedProofStored",
		"type": "event"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"name": "milestoneReleases",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "id",
				"type": "uint256"
			},
			{
				"internalType": "uint64",
				"name": "project_id",
				"type": "uint64"
			},
			{
				"internalType": "uint64",
				"name": "milestone_id",
				"type": "uint64"
			},
			{
				"internalType": "string",
				"name": "transfer_image_url",
				"type": "string"
			},
			{
				"internalType": "string",
				"name": "transfer_note",
				"type": "string"
			},
			{
				"internalType": "string",
				"name": "created_at",
				"type": "string"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"name": "verifiedProofs",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "id",
				"type": "uint256"
			},
			{
				"internalType": "uint64",
				"name": "project_id",
				"type": "uint64"
			},
			{
				"internalType": "uint64",
				"name": "milestone_id",
				"type": "uint64"
			},
			{
				"internalType": "string",
				"name": "transfer_receipt_url",
				"type": "string"
			},
			{
				"internalType": "string",
				"name": "proof_media_url",
				"type": "string"
			},
			{
				"internalType": "string",
				"name": "created_at",
				"type": "string"
			}
		],
		"stateMutability": "view",
		"type": "function"
	}
];