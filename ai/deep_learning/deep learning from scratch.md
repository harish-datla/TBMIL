## Math Foundations

It took very long for humans to invent the concept of counting, it was just 42k years earlier the [earliest evidence](https://en.wikipedia.org/wiki/History_of_ancient_numeral_systems) for counting was found, ancient humans just used to incise parallel marks on a baboons bone for each numerical, as in 1 is I, 2 is II, 3 is III, 4 is II|I, you get it.

it took approximately 40k years more at approximately 9th century BC(or much more as we cant put a finger on when the first human invented counting) to reach the famous roman numeral system and a thousand more years to get to present Arabic numerals(originating in India and adopted and spread to the world by Arabic mathematicians) and a thousand more for it to be fully adopted over the world. Its only just 500 years earlier, the world adopted the current Arabic numeral system.

**what is addition ?**
<details><summary>  </summary> <br>
	 FOR DISCRETE VALUES: There isn't an official definition per say, but as you already know addition is just counting the total count of discrete quantities from individual  collection of discrete quantities(like quantifiable items that can be called a single unit, a single person, a single book, a single paper, a single letter and so on),
	 FOR NON INTEGER REAL NUMBERS: Its finding the combined scale of the individual item. Its just an act of counting all the individual collection  of discrete quantities by putting them side by side
	  FOR CONTINUOUS(NON DISCRETE) VALUES**: addition for continuous values can be thought of adding quantifiably infinitesimal values of individual continuous values and then scaling it back to the integer level. TOO MUCH ? take an example -> 1.43 + 2.21 -> to make them discrete values you have to split them in to individual units of 0.01 , so 1.43 becomes 143 and 2.21 becomes 221, now you can add them 143 + 221 = 364, but we have to scale back remember ?   so to scale back 364 from 0.01 to 1 units, we have to scale back to 3.64, if you are too confused by the scaling and scale out read until multiplication and division and revisit here.
	</br>

**what is subtraction ?**
<details><summary>  </summary> <br>
	 FOR DISCRETE VALUES: Subtraction can be imagined as counting the absence of something, it can be then redefined in various ways.
	 it can also be thought of as opposite of addition. a - b , means , how much should i add to b discrete quantities to make the total a discrete quantities.  you can also imagine it as counting the absence of b discrete values in a pool of a discrete values.
	  FOR CONTINUOUS(NON DISCRETE) VALUES**:  The description above can be extended similar to the definition of adding non discrete values for subtraction(refer what is addition?)
	</br>
	
**What is multiplication ?**
<details><summary>  </summary> <br>
	 FOR DISCRETE VALUES: Multiplication is just repeated addition, if you have x number of equal discrete number ofcollections/groups(say y), how many discrete units are there in total? intuition says to count one by one or just add the number of discrete values x times, represented by x*y. You can also call this scaling, because we are scaling y , x times !
	  FOR CONTINUOUS(NON DISCRETE) VALUES**:  The description above can be extended similar to the definition of adding non discrete values for multiplication(refer what is addition?)
	</br>

**what is division ?**
<details><summary>  </summary> <br>
	 FOR DISCRETE VALUES: unlike addition, subtraction and multiplication, intuition for division is not that straightforward, fundamentally because multiplication is a little bit compicated than addition. you can think of it in 5 terms 1) inverse of multiplication 2) splitting apart, 3) scaling down 4) consolidating 5) repeated subtraction
	 1) Inverse of multiplication: how many x's make y (z = y/x) or how much we need to scale x to reach y?
	    at this point our brain finds it difficult to grasp this as our brain is not used to seeing things at this scale.
	2)  
	  FOR CONTINUOUS(NON DISCRETE) VALUES**:  The description above can be extended similar to the definition of adding non discrete values for multiplication(refer what is addition?)
	</br>

*Bonus Question*
**What is area of rectangle and why is it formulated the way it is?**
<details><summary>  </summary> <br>
	 Sure its length times breadth(l*b) , but why ? why did we specifically multiply ? you might say , we are multiplying the length l with breadth b, akin to sweeping the length across the breadth, forming the area. But arent there infinite lines of length l on the breadth line ? is l*b infinite ?
	 The answer is that we are counting the number of 1*1 squares inside the rectangle. there will be b no of 1*1 squares across the breadth and that forms a unit, and now we have l of these units across the length l, hence we multiply (l*b), so we abstract away area as the number of 1*1 squares in a given enclosed 2d-space. Any other definition will boil down to this fundamental idea(i think so, i have no proof)
	 
	Personal trivia : I didnt think about this until im like 18, much beyond my high school education one day, when i cant sleep, i was like why is rectangle area l*b, and i spent the next 5 days on it while simultaneously having existential crisis.
	</br>

#### A very trivial introduction to algebra
well, all you need to know about basic algebra is already covered(addition,subtraction,multiplication and division)barring couple of things. 1) the equality operator 2) variable naming
The equality operator
basically means what it says , everything put together to the left of the equality operator is equal to everything put together on the right side of it.
Variable Naming.








#### NLP with deep learning

##### Intro and word vectors.





#### Stanford Intro to LLM's
# What is a language model?

The classic definition of a language model (LM) is a probability distribution over sequences of tokens.

<details><summary> Explanation </summary> <br>
	- Self explanatory definition , if you have,  say n words in your vocabulary and you want to make a sentence with L words in a certain order, a language model will you give the probability of that sentence being legit syntactically(grammatically) and semantically(world knowledge, google for more)
	</br>

Suppose we have a vocabulary `V` of a set of tokens. A language model `p` assigns each sequence of tokens `x₁, …, x_L ∈ V` a probability (a number between 0 and 1):

$$p(x_1, \dots, x_L)$$

The probability intuitively tells us how "good" a sequence of tokens is.  For example, if the vocabulary is `V={ate,ball,cheese,mouse,the}` the language model might assign.
$$ p({the}, {mouse}, {ate}, {the}, {cheese}) = 0.02,$$
$$ p({the}, {cheese}, {ate}, {the}, {mouse}) = 0.01,$$
$$ p({mouse}, {the}, {the}, {cheese}, {ate}) = 0.0001,$$

although the  language model looks simple, the ability to assign meaningful probabilities to all sequences requires extraordinary and implicit linguistic abilities and world knowledge.

For example, the LM should assign the  *mouse the the cheese ate* a very low probability implicitly because its's ungrammatical(**syntactic knowledge**). The LM should assign *the mouse ate the cheese* higher probability than *the cheese ate the mouse* implicitly because of **World Knowledge:** both sentences are the same syntactically,  but they differ in semantic plausibility(semantic plausibility corresponds to world knowledge).

**Generation** -> a we can generate a sequence by the Language Model(defined above). A rudimentary way to do this would be to sample a sequence  $x_{1:L}$ from  language model $p$ with probability equal to $p(x_{1:L})$, denoted:
$$x_{1:L} \sim p.$$

How to do this computationally efficiently depends on the form of the language model $p$. IN practice, we do not generally sample directly from a language model both because of limitations of real language models and because we sometimes wish to obtain not an average sequence but something closer to the "best" sequence.

<details><summary> EXPLANATION </summary> <br>
	-Fancy of saying ->  given a language model(and/implicitly vocabulary) with above definition , we can generate as many sequences of words possible,  each permutation of the sequence has a certain probability associated with it. We cant just assume rudimentary high probability sequence to be the best, because well its probability and it might as well be an average sequence, what we want is the best sequence of tokens and we can achieve that by more refined techniques than just using a language model.
	</br>

**Auto regressive language models**.
A common way to write the joint distribution $p(x_{1:L})$ of a sequence $x_{1:L}$ is using the **chain rule of probability**:

$$ p(x_{1:L}) = p(x_1) p(x_2 \mid x_1) p(x_3 \mid x_1, x_2) \cdots p(x_L \mid x_{1:L-1}) = \prod_{i=1}^L p(x_i \mid x_{1:i-1}).$$ 
Example :
$$ \begin{align*} p({the},\,{mouse},\, {ate}, \,{the}, \,{cheese}) = \, & p(\,{the}) \\ & p(\,{mouse} \mid \,{the}) \\ & p(\,{ate} \mid \,{the}, \,{mouse}) \\ & p(\,{the} \mid \,{the}, \,{mouse}, \,{ate}) \\ & p(\,{cheese} \mid \,{the}, \,{mouse}, \,{ate}, \,{the}). \end{align*}$$

In particular, $p(x_i \mid x_{1:i-1})$ is a **conditional probability distribution** of the next token $x_i$ given the previous tokens $x_{1:i-1}$.

Well this is just probability distribution, nothing fancy.

an **auto regressive language model** is the one where each conditional distribution can be computed efficiently(for ex: using a feed forward neural network, more on that later or before)

**Generation**
Now to generate an entire sequence $x_{1:L}$ from autoregressive language model $p$, we sample one token at a time given the tokens generated so far.

$$
\text{for } i = 1,  \dots, L: \\ $$
$$ \hspace{1in} x_i \sim p(x_i \mid x_{1:i-1})^{1/T}
$$

where  $T \ge 0$ is a **temperature** parameter that controls how much randomness we want from the language model:

- $T = 0$: deterministically choose the most probable token $x_i$ at each position i
- $T = 1$: sample “normally” from the pure language model
- $T = \infty$: sample from a uniform distribution over the entire vocabulary $V$

However, if we just raise the probabilities to the power 1/T, the probability distribution may not sum to 1. We can fix this by re-normalizing the distribution. We call the normalized version $p_T(x_i \mid x_{1:i-1}) \propto p(x_i \mid x_{1:i-1})^{1/T}$ the **annealed** conditional probability distribution. For example:

$$ p({cheese}) = 0.4, \quad\quad\quad p({mouse}) = 0.6$$

$$p_{T=0.5}({cheese}) = 0.31, \quad\quad\quad p_{T=0.5}({mouse}) = 0.69$$

$$p_{T=0.2}({cheese}) = 0.12, \quad\quad\quad p_{T=0.2}({mouse}) = 0.88$$

$$p_{T=0}({cheese}) = 0, \quad\quad\quad p_{T=0}({mouse}) = 1 $$

_Aside_: Annealing is a reference to metallurgy, where hot materials are cooled gradually, and shows up in sampling and optimization algorithms such as simulated annealing.

_Technical note_: sampling iteratively with a temperature T parameter applied to each conditional distribution $p(x_i \mid x_{1:i-1})^{1/T}$ is not equivalent (except when T = 1) to sampling from the annealed distribution over length L sequences.

<details><summary> EXPLANATION </summary> <br>
	-Fancy of saying -> We can control the distribution of the language model according to our needs, for example T< 1 makes the distribution softer(i.e, softens the hard peaks in the probability distribution) , while T>1 makes the peaks more pronounced i.e, makes the distribution sharper. This will accordingly changes the model of our outputs. for example if we take 3 words as candidates for our current iteration, the next iteration can get us more interesting sentences with examples explained below.
	
	 Normalization formula is just calculating the new probabilities with the Temperature and normalizing it to 1. ie.,(a/(a+b), b/(a+b)).
	 
	 Hence as you can foresee , we can use T to get a language model that generates the best sequence texts.
	</br>

**Conditional generation**. More generally, we can perform conditional generation by specifying some prefix sequence $x_{1:i}$ (called a **prompt**) and sampling the rest $x_{i+1:L}$ (called the **completion**). For example, generating with T=0 produces

$$
\newcommand{\generate}[1]{\stackrel{#1}{\rightsquigarrow}}
\underbrace{\nl{the}, \nl{mouse}, \nl{ate}}_\text{prompt} \generate{T=0} \underbrace{\nl{the}, \nl{cheese}}_\text{completion}.$$

If we change the temperature to T = 1, we can get more variety , for example, **its house** and **my homework**.

As we’ll see shortly, conditional generation unlocks the ability for language models to solve a variety of tasks by simply changing the prompt.
<details><summary> EXPLANATION </summary> <br>
	-Fancy of saying ->We can control how many varied inputs we can get, if we set T=0, we can only get the most probable completion ie., the cheese, however just tweaking the T=1, can generate interesting inputs and we can choose these interesting inputs and comprehensively choose the best randomizer value(i.e., T)
	</br>

### Summary

- A language model is a probability distribution p over sequences $x_{1:L}$.
- Intuitively, a good language model should have linguistic capabilities and world knowledge.
- An autoregressive language model allows for efficient generation of a completion $x_{i+1:L}$ given a prompt $x_{1:i}$.
- The temperature can be used to control the amount of variability in generation.
	
 
