## AI-Generated Go Code for Go for Statistics


## AI-assisted programming (Using Github Copilot): 

In the copilot-suggestions branch, you will find the three files that were refactored using Github Copilot. You can compare and pull the changes from this branch and the main branch. It may say there is not much to compare, but they do have very similar features other than the changes made from the copilot. I would ask the chat questions and it would output a more efficient result. 

-Key Changes made: 

Main.go file: 
The copilot suggested increasing concurrency and avoid logging fatal errors to help improve efficiency of the test. We were able to run it at 1.9 ms. A waitgroup function was also added to wait for goroutines to finish first. 

Linear_regresson.go file: 
Added single point case handling to handle when there is only one data point. Pre-allocated slices were also added to stop repeated memory code. Overall, it improves the code’s efficiency. 

Linear_regression_test.file: I asked how to make this better and the copilot added multiple unit tests and a helper function of equalSlices to help compare equality. The benchmark stayed the same. However, when trying to test the files, the single point test failed and continued to fail the copilot was unable to fix that. However, being able to test valid input, mismatched input lengths, and edge cases helps better performance testing for the assignment which was helpful when using the copilot. 

Overall, I think Copilot is helpful to make code more efficient after writing it. However, it is important to check that it works correctly as it failed in the single point test still in the linear regression test file. I think it is also helpful to simplify code in loops and create suggestions when typing. 


## Selected LLM-Based Service

- # Step-by-Step Process:
- 1) Selecting the LLM Agent
     - I chose to use ChatGPT by OpenAI. The free plan was used here:  The service is available at https://chat.openai.com. The file CHATGPT-Dialogue explains the entire dialoge between me and the AI agent.
    
- 2) Generating Go Code for Statistics
     - Iterated through multiple prompts to refine the generated code and ensure correctness.
     - My plan was to basically copy the prompt from the assignment then I created different prompts from there
     - There were errors and I tried to fix it by asking AI how to fix it but the code was still invalid after multiple attempts
  
-  3) Stored dialoge
      - I stored the dialoge in the CHATGPT_dialogue pdf file in the repository
     
## Repository Files
CHATGPT_dialoge – Full conversation with ChatGPT, including prompts and responses.

copilot-suggestions – branch from github copilot with linearregression code and testing files
chatgpt-code -branch with chatgpt code and main and main test files from chatgpt. The copilot files show up in this one as well

## Summary of Experience

### 1. Automated Code Generation

This was helpful in suggesting copilot for what to comment or relative code when coding for developers. 

### 2. AI-Assisted Programming

Interacting with the Github Copilot AI allowed for rapid troubleshooting and optimization.

AI-assisted debugging and performance improvements saved time.

I like this the best becasue it helps developers create the code then make it cleaner once it is working. 

### 3. AI-Generated Code

While useful, AI-generated code often required manual verification.

ChatGPT provided  Go code for statistical analysis with minimal user intervention. The code was invalid which was not easy to figure out what went wrong.


## Reducing Workload and Recommendations
- At a startup this could help reduce responsibilities with debugging code and time. Junior developers could have higher developed code. 
- A mix of Ai and developers should be used
- It should be used more for debuggin and making code cleaner and more efficient
- Augmenting rather than replacing developers

## Materials Used
-Go

-ChatGPT (https://chatgpt.com/ )

-Github Copilot https://github.com/features/copilot?ef_id=_k_CjwKCAiAwaG9BhAREiwAdhv6Y9T51-Htvu15KFMnz2BH1_NC0asjLlMdUfGn8aWsgoP1A2zKGFTxkRoC_NgQAvD_BwE_k_&OCID=AIDcmmb150vbv1_SEM__k_CjwKCAiAwaG9BhAREiwAdhv6Y9T51-Htvu15KFMnz2BH1_NC0asjLlMdUfGn8aWsgoP1A2zKGFTxkRoC_NgQAvD_BwE_k_&gad_source=1&gclid=CjwKCAiAwaG9BhAREiwAdhv6Y9T51-Htvu15KFMnz2BH1_NC0asjLlMdUfGn8aWsgoP1A2zKGFTxkRoC_NgQAvD_BwE (I used this youtube channel to help: https://www.youtube.com/watch?v=z1ycDvspv8U )


