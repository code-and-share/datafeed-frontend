# Project Name
Code-and-share Frontend

# Preparation Notes
## Define the problem
I need a tool that fits a wide range of use cases to feed data into a Machine Learning Algorithm.  
  
I am not aiming at the complex models where input data can be very varied. Instead I am aiming at the easiest use cases, where variables A, B...N can have a finite number of different values, and the Machine Learning model can learn from that finite data. 
# Solution Brainstorming
## Solution Chosen
* Frontend written in HTML/CSS/JS, iterating with a dedicated Go Backend
* The Backend gets the data from a Postgres DB, and decides the groups depending on the responses received from the Frontend
* After all groups have had a selection, the Golang Backend sends those to a so-called “AI-backend”, which gives back the final result.
* The Golang backend sends the result to the frontend

# Related stakeholders
* Github user angelalonso
* Github user gamstc
# Competitors
TBD
# Objectives
## Specific
* I want to be able to dynamically present images/text…(from now on objects) to click on.
  * Show a predefined group :ok:
  * Make objects clickable from now on :ok:
  * Receive selection :ok:
  * Show a different site :ok:
  * Add first unit tests :ok:
  * Add at least one test per function :ok:
  * Show another predefined group :ok:
  * Save session_id cookie to identify all path :ok:
  * Add link to start new session :ok:
  * If new session, first site again :ok:
  * Receive second selection :ok:
  * Save selection orderly :ok:
  * Print both selections :ok:
    * ERROR: At this point my code just sends the same code to anyone going for the main page. I need to move this logic around:
    * Save current phase in cookies :ok:
    * Selection calls a different url :ok:
    * First page is redirected to phase/001
    * Array of results becomes array of objects [ session_id, results[]]
  * Make object IDs numbers or get a different value to send at Javascript
  * Return message after all phases including all selections
  * ...
* Those objects are presented in groups of 4, but could be more or less than that. Important is that the user has as little bias as possible.
* What the user selects on groups (1..n) defines the objects presented on group n+1. Group 1 should also be random.
* The amount of groups to be shown depends on the user selection.
* After all groups have had a selection made, the user is presented with a result.

# Possible trade-offs between Quality – Time – Cost
TBD

# Estimate project resources
TBD






