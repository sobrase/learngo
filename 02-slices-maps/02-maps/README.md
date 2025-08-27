# Exercise 4: Social Network User Management System

## 🎯 **Objective**
Create a social network user management system that demonstrates Go maps, their operations, and working with structured data.

## 🎨 **Goal: Print a Social Network Dashboard**

You will create a program that manages users, their connections, posts, and generates a beautiful social network dashboard.

## 📋 **Requirements**
1. Create and initialize maps in different ways:
   - Using `make()`
   - Using literal syntax
   - Using `map[T]T{}` syntax
2. Perform map operations:
   - Adding key-value pairs
   - Updating values
   - Deleting keys
   - Checking if key exists
3. Work with nested maps and complex data structures
4. Use maps for counting and grouping data
5. Iterate over maps using `range`

## 🎯 **Specific Tasks**

### **Task 1: User Profile Management**
Create maps to store user information:
- User profiles (ID -> Profile data)
- User connections (ID -> []Friend IDs)
- User posts (ID -> []Post data)
- User statistics (ID -> Stats data)

### **Task 2: Social Network Operations**
Implement functions for:
- `addUser(id string, profile map[string]interface{})` - Add new user
- `addConnection(userID, friendID string)` - Connect users
- `addPost(userID string, post map[string]interface{})` - Add user post
- `getUserStats(userID string)` - Calculate user statistics
- `findMutualFriends(user1, user2 string)` - Find common friends

### **Task 3: Data Analysis**
Create functions for:
- `getMostActiveUsers()` - Find users with most posts
- `getPopularUsers()` - Find users with most connections
- `getTrendingTopics()` - Find most mentioned topics
- `getUserRecommendations(userID string)` - Suggest new friends

### **Task 4: Print Social Network Dashboard**
Generate a beautiful dashboard that looks like this:

```
╔══════════════════════════════════════════════════════════════╗
║                 SOCIAL NETWORK DASHBOARD                     ║
╠══════════════════════════════════════════════════════════════╣
║                                                              ║
║  📊 NETWORK STATISTICS:                                      ║
║  • Total Users: 5                                            ║
║  • Total Connections: 8                                      ║
║  • Total Posts: 12                                           ║
║  • Average Posts per User: 2.4                               ║
║                                                              ║
╠══════════════════════════════════════════════════════════════╣
║  👥 USER PROFILES:                                           ║
║  ┌─────────────┬─────────────┬─────────────┬─────────────┐   ║
║  │ User ID     │ Name        │ Age         │ Location    │   ║
║  ├─────────────┼─────────────┼─────────────┼─────────────┤   ║
║  │ U001        │ Alice       │ 25          │ New York    │   ║
║  │ U002        │ Bob         │ 30          │ Los Angeles │   ║
║  │ U003        │ Charlie     │ 28          │ Chicago     │   ║
║  │ U004        │ Diana       │ 22          │ Miami       │   ║
║  │ U005        │ Eve         │ 27          │ Seattle     │   ║
║  └─────────────┴─────────────┴─────────────┴─────────────┘   ║
║                                                              ║
╠══════════════════════════════════════════════════════════════╣
║  🔗 CONNECTIONS NETWORK:                                     ║
║                                                              ║
║  Alice (U001) ──── Bob (U002) ──── Charlie (U003)           ║
║     │              │              │                          ║
║     └─── Diana (U004) ──── Eve (U005)                        ║
║                                                              ║
║  Connection Details:                                          ║
║  • Alice ↔ Bob (Mutual)                                      ║
║  • Bob ↔ Charlie (Mutual)                                    ║
║  • Alice ↔ Diana (Mutual)                                    ║
║  • Charlie ↔ Eve (Mutual)                                    ║
║  • Diana ↔ Eve (Mutual)                                      ║
║                                                              ║
╠══════════════════════════════════════════════════════════════╣
║  📝 RECENT POSTS:                                            ║
║  ┌─────────────┬─────────────┬─────────────┬─────────────┐   ║
║  │ User        │ Post        │ Likes       │ Comments    │   ║
║  ├─────────────┼─────────────┼─────────────┼─────────────┤   ║
║  │ Alice       │ "Hello!"    │ 5           │ 2           │   ║
║  │ Bob         │ "Great day" │ 8           │ 3           │   ║
║  │ Charlie     │ "Coding..." │ 12          │ 5           │   ║
║  │ Diana       │ "Weekend!"  │ 3           │ 1           │   ║
║  │ Eve         │ "New job!"  │ 15          │ 7           │   ║
║  └─────────────┴─────────────┴─────────────┴─────────────┘   ║
║                                                              ║
╠══════════════════════════════════════════════════════════════╣
║  🏆 LEADERBOARD:                                             ║
║                                                              ║
║  Most Active Users:                                          ║
║  1. Eve (U005) - 4 posts, 15 total likes                    ║
║  2. Charlie (U003) - 3 posts, 12 total likes                ║
║  3. Bob (U002) - 2 posts, 8 total likes                     ║
║                                                              ║
║  Most Connected Users:                                       ║
║  1. Bob (U002) - 3 connections                               ║
║  2. Alice (U001) - 2 connections                             ║
║  3. Charlie (U003) - 2 connections                           ║
║                                                              ║
╠══════════════════════════════════════════════════════════════╣
║  💡 RECOMMENDATIONS:                                         ║
║                                                              ║
║  For Alice (U001):                                           ║
║  • Connect with Charlie (2 mutual friends)                   ║
║  • Connect with Eve (1 mutual friend)                        ║
║                                                              ║
║  Trending Topics:                                            ║
║  • #weekend (3 mentions)                                     ║
║  • #coding (2 mentions)                                      ║
║  • #newjob (1 mention)                                       ║
║                                                              ║
╚══════════════════════════════════════════════════════════════╝
```

## 🎯 **Expected Output**
Your program should produce:
- User profiles and connections
- Social network visualization
- Recent posts and statistics
- Leaderboards and recommendations
- Formatted dashboard display

## 💡 **Hints**
- Use nested maps for complex data structures
- Use `range` to iterate over maps
- Use `delete()` to remove connections
- Use `make()` to initialize maps with capacity
- Use `fmt.Printf` with width specifiers for alignment
- Use Unicode symbols for visual appeal

## 🚀 **Bonus Challenges**
1. Add user authentication system
2. Implement post privacy settings
3. Add real-time notifications
4. Create user activity timeline
5. Add photo/video post support
6. Implement search functionality

## 🎨 **Learning Outcomes**
By completing this exercise, you will:
- Master map creation and manipulation
- Understand nested data structures
- Learn data analysis and statistics
- Create complex data relationships
- Practice real-world social media concepts

This exercise transforms map operations into an engaging social network management system! 🌐
