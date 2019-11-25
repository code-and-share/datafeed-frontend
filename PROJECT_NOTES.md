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
* I want to be able to get files from remote (set source as variable) :ok:
  * Make the url an environment variable :ok:
* I want to be able to dynamically present images/text…(from now on objects) to click on. :x:
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
  * If session cookie does not exist, new session_id and first site again :ok:
  * If phase cookie does not exist, phase is 1 so first site again :ok:
  * Receive second selection :ok:
  * Save selection orderly :ok:
  * Save current phase in cookies :ok:
  * Array of results becomes array of objects [ session_id, results[]] :ok:
  * Save all selections by session_id :ok:
  * Print only selections for this session_id :ok:
  * Build a standard number of phases and a results page at the end :ok:
  * Make all tests pass again :ok:
  * Make object IDs numbers or get a different value to send at Javascript :ok:
  * Return message after all phases including all selections :ok:
  * Add a database that stores the objects to show :ok:
  * Use the database instead :ok:
  * Take the full path from the object -> select object content from the JSON lists :ok:
  * Get the use of a DB tested :ok:
  * Get different tests when the DB is unavailable
  * Get config from ENV Variables :ok:
  * Containerize the app :ok:
  * Docker compose to also get DB up :ok:
  * Move to postgresql :ok:
  * Show proper error when DB is not available :x:
  * Show proper error when an object cannot be found (index does not exist) :x:
  * Add a backend service that can add paths, phases and objects :x:
  * Make the backend service also chooses the default path to follow by the main program :x:
  * Create a table of results :x:
  * Make Webdata receive a flexible amount of objects :x:
  * Adapt Css to the amount of objects received :x:
  * Beautify result :x:
  * ...
* Those objects are presented in groups of 4, but could be more or less than that. Important is that the user has as little bias as possible.
* What the user selects on groups (1..n) defines the objects presented on group n+1. Group 1 should also be random.
* The amount of groups to be shown depends on the user selection.
* After all groups have had a selection made, the user is presented with a result.

# Possible trade-offs between Quality – Time – Cost
TBD

# Estimate project resources
TBD






