# Import the package and module
from awadb.client import Awa

# Initialize awadb client
client = Awa()

# Add dict with vector to table 'example1'
client.add("example1", {'name':'david', 'feature':[1.3, 2.5, 1.9]})
client.add("example1", {'name':'jim', 'feature':[1.1, 1.4, 2.3]})

# Search
results = client.search("example1", [1.0, 2.0, 3.0])

# Output results
print(results)