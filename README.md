# Assignment6
This is creating linear regression model for CSV 'Boston". Creating linear regression models for concurrency and non-concurrency and comnpare thier execution time by running the each program 100 times

The execution time of non-currency program is longer than concureency's. The total execution time for 100 runs in nonconcurrency program takes 166.4593ms and the average per run is 1.664593ms.The total execution time for 100 runs in concurrency program tkaes 118.7686ms and the average per run is 1.187686ms.

# Advantage of Concureency
Increased Efficiency: Concurrency enables multiple processes to be executed simultaneously, which can significantly reduce the time required for large computations, such as training and testing multiple machine learning models.
Reduced Idle Time: By executing tasks in parallel, the overall idle time of the system is reduced, as the system can handle other tasks while waiting for I/O operations or other blocking events.
Complexity Management: Concurrency can simplify the management of complex, independent tasks by running them as separate concurrent processes.

# Advise for management
Concurrency in Go allows tasks to be executed concurrently, leveraging multiple CPU cores effectively. Company should use more concurrency and this will increase company's efficiency.
