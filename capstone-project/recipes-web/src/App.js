import React, { useState, useEffect } from 'react';
import './App.css';

function App() {
    const [recipes, setRecipes] = useState([]);

    const [searchText, setSearchText] = useState("");
    const [searchResults, setSearchResults] = useState([]);
    const [searchedRecipe, setSearchedRecipe] = useState(null);
    const [searchError, setSearchError] = useState("");
    const [isSearching, setIsSearching] = useState(false);
    const [isLoadingDetail, setIsLoadingDetail] = useState(false);

    useEffect(() => {
        fetch("http://localhost:8088/recipes")
            .then(res => res.json())
            .then(data => {
                setRecipes(data);
            })
            .catch(err => console.error("Error fetching recipes:", err));
    }, []);

    const handleSearch = async () => {
        const query = searchText.trim();
        if (!query) return;

        setIsSearching(true);
        setSearchError("");
        setSearchedRecipe(null);

        try {
            const response = await fetch(`http://localhost:8088/recipes/search?q=${encodeURIComponent(query)}`);
            if (!response.ok) {
                throw new Error("Search failed");
            }

            const data = await response.json();
            const results = Array.isArray(data) ? data : [];
            setSearchResults(results);
            if (results.length === 0) {
                setSearchError("No recipes found for this search.");
            }
        } catch (err) {
            setSearchError(err.message || "Search failed");
            setSearchResults([]);
        } finally {
            setIsSearching(false);
        }
    };

    const handleRecipeIdClick = async (id) => {
        setIsLoadingDetail(true);
        setSearchError("");

        try {
            const response = await fetch(`http://localhost:8088/recipe/${id}`);
            if (!response.ok) {
                throw new Error("Recipe not found");
            }

            const data = await response.json();
            setSearchedRecipe(data);
        } catch (err) {
            setSearchError(err.message || "Failed to fetch recipe details");
            setSearchedRecipe(null);
        } finally {
            setIsLoadingDetail(false);
        }
    };

    const clearSearch = () => {
        setSearchText("");
        setSearchResults([]);
        setSearchedRecipe(null);
        setSearchError("");
    };

    const backToResults = () => {
        setSearchedRecipe(null);
    };

    const getRecipeImage = (id) => {
        return `https://picsum.photos/seed/${id}/400/300`;
    };

    return (
        <div className="App">
            <header>
                <h1>Recipe Platform</h1>
            </header>

            <div className="search-container">
                <input
                    className="search-input"
                    type="text"
                    value={searchText}
                    onChange={(e) => setSearchText(e.target.value)}
                    onKeyDown={(e) => e.key === "Enter" && handleSearch()}
                    placeholder="Search recipes by name or tag..."
                />
                <button className="btn btn-primary" onClick={handleSearch} disabled={isSearching}>
                    {isSearching ? "Searching..." : "Search"}
                </button>
                <button className="btn btn-secondary" onClick={clearSearch}>Clear</button>
            </div>

            <p className="search-hint">Try searches like: paneer, biryani, indian, vegetarian</p>

            {searchError && <p className="error-message">{searchError}</p>}

            {searchedRecipe ? (
                <div>
                    <div className="results-actions">
                        <h2 className="section-title">Recipe Details</h2>
                        <button className="btn btn-primary" onClick={backToResults}>Back to Results</button>
                    </div>
                    <div className="recipe-grid" style={{ gridTemplateColumns: 'minmax(300px, 500px)' }}>
                        <div className="recipe-card">
                            <img
                                src={searchedRecipe.imageUrl ? `http://localhost:8088/${searchedRecipe.imageUrl}` : getRecipeImage(searchedRecipe.id)}
                                alt={searchedRecipe.name}
                                className="recipe-image"
                            />
                            <div className="recipe-content">
                                <h3 className="recipe-title">{searchedRecipe.name}</h3>
                                <p className="recipe-id">ID: {searchedRecipe.id}</p>
                                <div className="recipe-details">
                                    <strong>Ingredients:</strong>
                                    {searchedRecipe.ingredients && searchedRecipe.ingredients.length > 0 ? (
                                        <ul style={{ paddingLeft: "20px", marginTop: "5px" }}>
                                            {searchedRecipe.ingredients.map((ingredient, index) => (
                                                <li key={index}>{ingredient}</li>
                                            ))}
                                        </ul>
                                    ) : (
                                        <p>None</p>
                                    )}
                                </div>

                                <div className="recipe-details">
                                    <strong>Instructions:</strong>
                                    {searchedRecipe.instructions && searchedRecipe.instructions.length > 0 ? (
                                        <ol style={{ paddingLeft: "20px", marginTop: "5px" }}>
                                            {searchedRecipe.instructions.map((instruction, index) => (
                                                <li key={index}>{instruction}</li>
                                            ))}
                                        </ol>
                                    ) : (
                                        <p>None</p>
                                    )}
                                </div>
                                <div className="recipe-tags">
                                    {searchedRecipe.tags ? searchedRecipe.tags.map(tag => (
                                        <span key={tag} className="tag">{tag}</span>
                                    )) : <span className="tag">No tags</span>}
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            ) : searchResults.length > 0 ? (
                <div>
                    <div className="results-actions">
                        <h2 className="section-title">Search Results</h2>
                        <span className="search-meta">{searchResults.length} recipe(s) found</span>
                    </div>
                    <div className="recipe-grid">
                        {searchResults.map((recipe) => (
                            <div key={recipe.id} className="recipe-card">
                                <img
                                    src={recipe.imageUrl ? `http://localhost:8088/${recipe.imageUrl}` : getRecipeImage(recipe.id)}
                                    alt={recipe.name}
                                    className="recipe-image"
                                />
                                <div className="recipe-content">
                                    <h3 className="recipe-title">{recipe.name}</h3>
                                    <p className="recipe-id">
                                        ID:{" "}
                                        <button
                                            type="button"
                                            className="recipe-id-link"
                                            onClick={() => handleRecipeIdClick(recipe.id)}
                                            disabled={isLoadingDetail}
                                        >
                                            {recipe.id}
                                        </button>
                                    </p>
                                    <p className="recipe-details">Click the highlighted ID to open full recipe details.</p>
                                    <div className="recipe-tags">
                                        {recipe.tags ? recipe.tags.map(tag => (
                                            <span key={tag} className="tag">{tag}</span>
                                        )) : <span className="tag">No tags</span>}
                                    </div>
                                </div>
                            </div>
                        ))}
                    </div>
                </div>
            ) : (
                <div>
                    <h2 className="section-title">All Recipes</h2>
                    <div className="recipe-grid">
                        {recipes && recipes.map((recipe) => (
                            <div key={recipe.id} className="recipe-card">
                                <img
                                    src={recipe.imageUrl ? `http://localhost:8088/${recipe.imageUrl}` : getRecipeImage(recipe.id)}
                                    alt={recipe.name}
                                    className="recipe-image"
                                />
                                <div className="recipe-content">
                                    <h3 className="recipe-title">{recipe.name}</h3>
                                    <p className="recipe-id">ID: {recipe.id}</p>
                                    <div className="recipe-details">
                                        <strong>Ingredients:</strong>
                                        {recipe.ingredients && recipe.ingredients.length > 0 ? (
                                            <ul style={{ paddingLeft: "20px", marginTop: "5px" }}>
                                                {recipe.ingredients.map((ingredient, index) => (
                                                    <li key={index}>{ingredient}</li>
                                                ))}
                                            </ul>
                                        ) : (
                                            <p>None</p>
                                        )}
                                    </div>

                                    <div className="recipe-details">
                                        <strong>Instructions:</strong>
                                        {recipe.instructions && recipe.instructions.length > 0 ? (
                                            <ol style={{ paddingLeft: "20px", marginTop: "5px" }}>
                                                {recipe.instructions.map((instruction, index) => (
                                                    <li key={index}>{instruction}</li>
                                                ))}
                                            </ol>
                                        ) : (
                                            <p>None</p>
                                        )}
                                    </div>
                                    <div className="recipe-tags">
                                        {recipe.tags ? recipe.tags.map(tag => (
                                            <span key={tag} className="tag">{tag}</span>
                                        )) : <span className="tag">No tags</span>}
                                    </div>
                                </div>
                            </div>
                        ))}
                    </div>
                </div>
            )}
        </div>
    );
}
export default App;